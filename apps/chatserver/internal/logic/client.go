package logic

import (
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"jt-chat/common/constant"
	protocol "jt-chat/common/pb"
)

type Client struct {
	Conn      *websocket.Conn
	Uid       string
	Send      chan []byte
	StopWrite chan bool
	StopRead  chan bool
	Ctx       context.Context
	Server    *Server
}

func (c *Client) Write() {
	for {
		select {
		case <-c.StopWrite:
			return
		case message := <-c.Send:
			err := c.Conn.WriteMessage(websocket.BinaryMessage, message)
			if err != nil {
				logx.WithContext(c.Ctx).Error(errors.Wrapf(err, "用户发送消息"))
			}
		}
	}
}

func (c *Client) Read() {
	for {
		select {
		case <-c.StopRead:
			return
		default:
			c.Conn.PongHandler()
			_, message, err := c.Conn.ReadMessage()
			if err != nil {
				// websocket断开连接
				c.Server.Cancellation <- c
				logx.WithContext(c.Ctx).Error(errors.Wrapf(err, "用户接收消息"))
				continue
			}
			msg := &protocol.Message{}
			err = proto.Unmarshal(message, msg)
			if err != nil {
				logx.WithContext(c.Ctx).Error(errors.Wrapf(err, "消息解码"))
				continue
			}
			if msg.TransportType == constant.TransportTypeHeartBeat {
				pong := &protocol.Message{
					Content:       constant.MsgPong,
					TransportType: constant.TransportTypeHeartBeat,
				}
				pongBytes, err := proto.Marshal(pong)
				if err != nil {
					logx.WithContext(c.Ctx).Error(errors.Wrapf(err, "消息编码"))
					continue
				}
				err = c.Conn.WriteMessage(websocket.BinaryMessage, pongBytes)
				if err != nil {
					logx.WithContext(c.Ctx).Error(errors.Wrapf(err, "发送心跳信息"))
				}
			} else {
				c.Server.Send <- message
			}
		}
	}
}
