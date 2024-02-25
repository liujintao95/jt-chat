package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"jt-chat/apps/message/api/internal/config"
	"jt-chat/apps/message/rpc/message"
)

type ServiceContext struct {
	Config config.Config
	MsgRpc message.MessageZrpcClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		MsgRpc: message.NewMessageZrpcClient(zrpc.MustNewClient(c.UserRpcConf)),
	}
}
