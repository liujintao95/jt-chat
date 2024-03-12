package service

import (
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"jt-chat/common/constant"
	protocol "jt-chat/common/pb"
	"net/http"
	"time"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Client struct {
	Conn      *websocket.Conn
	Uid       string
	Send      chan []byte
	StopWrite chan bool
	StopRead  chan bool
	Ctx       context.Context
	Hub       *Hub
}

func (c *Client) Write() {
	ticker := time.NewTicker(constant.PingPeriod)
	defer func() {
		ticker.Stop()
		_ = c.Conn.Close()
	}()
	for {
		select {
		case <-ticker.C:
			_ = c.Conn.SetWriteDeadline(time.Now().Add(constant.WriteWait))
			err := c.Conn.WriteMessage(websocket.PingMessage, nil)
			if err != nil {
				return
			}
		case message, ok := <-c.Send:
			if !ok {
				_ = c.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			_ = c.Conn.SetWriteDeadline(time.Now().Add(constant.WriteWait))
			err := c.Conn.WriteMessage(websocket.BinaryMessage, message)
			if err != nil {
				logx.WithContext(c.Ctx).Error(errors.Wrapf(err, "用户发送消息"))
				return
			}
		}
	}
}

func (c *Client) Read() {
	defer func() {
		c.Hub.Cancellation <- c
		_ = c.Conn.Close()
	}()
	for {
		_, message, err := c.Conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				logx.WithContext(c.Ctx).Error(errors.Wrapf(err, "用户接收消息"))
			}
			break
		}
		msg := &protocol.MessageForm{}
		err = proto.Unmarshal(message, msg)
		if err != nil {
			logx.WithContext(c.Ctx).Error(errors.Wrapf(err, "消息解码"))
			continue
		}
		c.Hub.Send <- message
	}
}

func ServeWs(ctx context.Context, hub *Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logx.WithContext(ctx).Error(errors.Wrapf(err, "建立websocket连接"))
		return
	}
	client := &Client{
		Conn:      conn,
		Uid:       r.URL.Query().Get(constant.UidKey),
		Send:      make(chan []byte),
		StopWrite: make(chan bool),
		StopRead:  make(chan bool),
		Ctx:       ctx,
		Hub:       hub,
	}
	client.Hub.registerClient(client)
	go client.Read()
	go client.Write()
}
