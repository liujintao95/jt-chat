package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"jt-chat/apps/user/api/internal/config"
	"jt-chat/apps/user/rpc/user"
)

type ServiceContext struct {
	Config  config.Config
	UserRpc user.UserZrpcClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: user.NewUserZrpcClient(zrpc.MustNewClient(c.UserRpcConf)),
	}
}
