package service

import (
	"context"
	"github.com/golang/protobuf/proto"
	"github.com/gorilla/websocket"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"jt-chat/apps/chatserver/internal/svc"
	messagePb "jt-chat/apps/message/rpc/pb"
	userPb "jt-chat/apps/user/rpc/pb"
	"jt-chat/common/constant"
	"jt-chat/common/ctxdata"
	protocol "jt-chat/common/pb"
	"sync"
)

type Hub struct {
	Mutex        *sync.Mutex
	Clients      map[string]*Client
	Send         chan []byte
	Register     chan *Client
	Cancellation chan *Client
	Ctx          context.Context
	SvcCtx       *svc.ServiceContext
	Logger       logx.Logger
}

func NewSocketHub(ctx context.Context, svcCtx *svc.ServiceContext) *Hub {
	return &Hub{
		Mutex:        &sync.Mutex{},
		Clients:      make(map[string]*Client),
		Send:         make(chan []byte),
		Register:     make(chan *Client),
		Cancellation: make(chan *Client),
		Ctx:          ctx,
		SvcCtx:       svcCtx,
		Logger:       logx.WithContext(ctx),
	}
}

func (h *Hub) Start() {
	logx.WithContext(h.Ctx).Info("聊天服务器已启动")
	for {
		select {
		case client := <-h.Register:
			h.registerClient(client)
			logx.WithContext(h.Ctx).Infof("%s登录聊天服务器", client.Uid)
		case client := <-h.Cancellation:
			h.cancellationClient(client)
			logx.WithContext(h.Ctx).Infof("%s登出聊天服务器", client.Uid)
		case message := <-h.Send:
			// 信息发送
			msg := &protocol.MessageForm{}
			err := proto.Unmarshal(message, msg)
			if err != nil {
				logx.WithContext(h.Ctx).Error(errors.Wrapf(err, "消息解码"))
				continue
			}
			_, exits := h.Clients[msg.From]
			if !exits {
				continue
			}
			h.saveMsg(msg)
			if msg.ToType == constant.ChatSingle {
				h.sendSingleMsg(msg)
			} else {
				h.sendGroupMsg(msg)
			}
		}
	}
}

func (h *Hub) registerClient(client *Client) {
	if _, ok := h.Clients[client.Uid]; ok {
		h.Cancellation <- client
	}
	h.Clients[client.Uid] = client
	msg := &protocol.MessageForm{
		From:          constant.AdminUid,
		To:            client.Uid,
		ToType:        constant.ChatSingle,
		Content:       constant.MsgWelcome,
		ContentType:   constant.ContentTypeText,
		TransportType: constant.TransportTypeNormal,
	}
	msgBytes, err := proto.Marshal(msg)
	if err != nil {
		logx.WithContext(h.Ctx).Error(errors.Wrapf(err, "消息编码"))
		return
	}
	err = client.Conn.WriteMessage(websocket.BinaryMessage, msgBytes)
	if err != nil {
		logx.WithContext(h.Ctx).Error(errors.Wrapf(err, "服务器发送消息"))
	}
}

func (h *Hub) cancellationClient(client *Client) {
	if _, ok := h.Clients[client.Uid]; ok {
		close(client.Send)
		delete(h.Clients, client.Uid)
	}
}

func (h *Hub) sendSingleMsg(msg *protocol.MessageForm) {
	client, ok := h.Clients[msg.To]
	if ok {
		msgByte, err := proto.Marshal(msg)
		if err == nil {
			client.Send <- msgByte
		}
	}
}

func (h *Hub) sendGroupMsg(msg *protocol.MessageForm) {
	var (
		groupUsers *userPb.GetGroupUserListOut
		err        error
	)
	groupUsers, err = h.SvcCtx.UserRpc.GetGroupUserList(
		ctxdata.GenerateCtx(h.Ctx, msg.From),
		&userPb.GetGroupUserListIn{
			Gid:  msg.To,
			Page: 1,
			Size: constant.MaxGroupSize,
		},
	)
	if err != nil {
		logx.WithContext(h.Ctx).Error(errors.Wrapf(err, "获取群人员信息"))
		return
	}
	for _, userInfo := range groupUsers.GroupUserList {
		if userInfo.Uid == msg.From {
			continue
		}
		sendMsg := msg
		sendMsg.From = msg.To
		sendMsg.To = userInfo.Uid
		client, ok := h.Clients[msg.To]
		if !ok {
			continue
		}
		msgByte, err := proto.Marshal(sendMsg)
		if err == nil {
			client.Send <- msgByte
		}
	}
}

func (h *Hub) saveMsg(msg *protocol.MessageForm) {
	var (
		newMsg     *messagePb.CreateMsgIn
		uploadResp *messagePb.UploadFileOut
		err        error
	)
	newMsg = &messagePb.CreateMsgIn{
		TransportType: msg.TransportType,
		From:          msg.From,
		To:            msg.To,
		ToType:        msg.ToType,
		Content:       msg.Content,
		ContentType:   msg.ContentType,
		FileExt:       msg.FileExt,
	}
	if len(msg.File) != 0 {
		uploadResp, err = h.SvcCtx.MsgRpc.UploadFile(h.Ctx, &messagePb.UploadFileIn{
			Name: msg.FileName,
			Ext:  msg.FileExt,
			Data: msg.File,
		})
		if err != nil {
			logx.WithContext(h.Ctx).Error(errors.Wrapf(err, "hub保存消息中的文件"))
			return
		}
		newMsg.FilePath = uploadResp.FilePath
	}
	_, err = h.SvcCtx.MsgRpc.CreateMsg(h.Ctx, newMsg)
	if err != nil {
		logx.WithContext(h.Ctx).Error(errors.Wrapf(err, "hub保存文件信息"))
		return
	}
}

func (h *Hub) Stop() {
	for _, client := range h.Clients {
		h.Cancellation <- client
	}
	_ = logx.Close()
}
