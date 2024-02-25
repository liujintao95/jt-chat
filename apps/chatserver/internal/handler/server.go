package handler

import (
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/zrpc"
	"jt-chat/apps/message/rpc/message"
	"jt-chat/apps/user/rpc/pb"
	"jt-chat/apps/user/rpc/user"
	"jt-chat/common/constant"
	"jt-chat/common/ctxdata"
	protocol "jt-chat/common/pb"
	"sync"
)

//var ChatServer = NewSocketServer()

type Server struct {
	Mutex        *sync.Mutex
	Clients      map[string]*Client
	Send         chan []byte
	Register     chan *Client
	Cancellation chan *Client
	Ctx          context.Context
	UserRpc      user.UserZrpcClient
	MsgRpc       message.MessageZrpcClient
	Logger       logx.Logger
}

func NewSocketServer() *Server {
	ctx := context.Background()
	return &Server{
		Mutex:        &sync.Mutex{},
		Clients:      make(map[string]*Client),
		Send:         make(chan []byte),
		Register:     make(chan *Client),
		Cancellation: make(chan *Client),
		Ctx:          context.Background(),
		UserRpc: user.NewUserZrpcClient(zrpc.MustNewClient(zrpc.RpcClientConf{
			Endpoints: []string{"127.0.0.1:2004"},
			NonBlock:  true,
		})),
		MsgRpc: message.NewMessageZrpcClient(zrpc.MustNewClient(zrpc.RpcClientConf{
			Endpoints: []string{"127.0.0.1:2005"},
			NonBlock:  true,
		})),
		Logger: logx.WithContext(ctx),
	}
}

func (s *Server) Start() {
	logx.WithContext(s.Ctx).Info("聊天服务器已启动")
	for {
		select {
		case client := <-s.Register:
			s.registerClient(client)
			logx.WithContext(s.Ctx).Infof("%s登录聊天服务器", client.Uid)
		case client := <-s.Cancellation:
			s.cancellationClient(client)
			logx.WithContext(s.Ctx).Infof("%s登出聊天服务器", client.Uid)
		case message := <-s.Send:
			// 信息发送
			msg := &protocol.Message{}
			err := proto.Unmarshal(message, msg)
			if err != nil {
				logx.WithContext(s.Ctx).Error(errors.Wrapf(err, "消息解码"))
				continue
			}
			_, exits := s.Clients[msg.From]
			if !exits {
				continue
			}
			s.saveMsg(msg)
			if msg.ToType == constant.ChatSingle {
				s.sendSingleMsg(msg)
			} else {
				s.sendGroupMsg(msg)
			}
		}
	}
}

func (s *Server) registerClient(client *Client) {
	go client.Read()
	go client.Write()
	if _, ok := s.Clients[client.Uid]; ok {
		s.Cancellation <- client
	}
	s.Clients[client.Uid] = client
	msg := &protocol.Message{
		From:          constant.AdminUid,
		To:            client.Uid,
		ToType:        constant.ChatSingle,
		Content:       constant.MsgWelcome,
		ContentType:   constant.ContentTypeText,
		TransportType: constant.TransportTypeNormal,
	}
	msgBytes, err := proto.Marshal(msg)
	if err != nil {
		logx.WithContext(s.Ctx).Error(errors.Wrapf(err, "消息编码"))
		return
	}
	err = client.Conn.WriteMessage(websocket.BinaryMessage, msgBytes)
	if err != nil {
		logx.WithContext(s.Ctx).Error(errors.Wrapf(err, "服务器发送消息"))
	}
}

func (s *Server) cancellationClient(client *Client) {
	if _, ok := s.Clients[client.Uid]; ok {
		_ = client.Conn.Close()
		client.StopRead <- true
		client.StopWrite <- true
		close(client.StopRead)
		close(client.StopWrite)
		close(client.Send)
		delete(s.Clients, client.Uid)
	}
}

func (s *Server) sendSingleMsg(msg *protocol.Message) {
	client, ok := s.Clients[msg.To]
	if ok {
		msgByte, err := proto.Marshal(msg)
		if err == nil {
			client.Send <- msgByte
		}
	}
}

func (s *Server) sendGroupMsg(msg *protocol.Message) {
	var (
		groupUsers *pb.GetGroupUserListOut
		err        error
	)
	groupUsers, err = s.UserRpc.GetGroupUserList(
		ctxdata.GenerateCtx(s.Ctx, msg.From),
		&pb.GetGroupUserListIn{
			Gid:  msg.To,
			Page: 1,
			Size: constant.MaxGroupSize,
		},
	)
	if err != nil {
		logx.WithContext(s.Ctx).Error(errors.Wrapf(err, "获取群人员信息"))
		return
	}
	for _, userInfo := range groupUsers.GroupUserList {
		if userInfo.Uid == msg.From {
			continue
		}
		sendMsg := msg
		sendMsg.From = msg.To
		sendMsg.To = userInfo.Uid
		client, ok := s.Clients[msg.To]
		if !ok {
			continue
		}
		msgByte, err := proto.Marshal(sendMsg)
		if err == nil {
			client.Send <- msgByte
		}
	}
}

func (s *Server) saveMsg(msg *protocol.Message) {
	// todo
}

func (s *Server) Stop() {
	for _, client := range s.Clients {
		s.Cancellation <- client
	}
	_ = logx.Close()
}
