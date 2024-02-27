package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"jt-chat/apps/chatserver/internal/config"
	"jt-chat/apps/message/rpc/message"
	"jt-chat/apps/user/rpc/user"
)

type ServiceContext struct {
	Config config.Config

	UserRpc user.UserZrpcClient
	MsgRpc  message.MessageZrpcClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: user.NewUserZrpcClient(zrpc.MustNewClient(c.UserRpcConf)),
		MsgRpc:  message.NewMessageZrpcClient(zrpc.MustNewClient(c.MsgRpcConf)),
	}
}
