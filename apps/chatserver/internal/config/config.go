package config

import (
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	service.ServiceConf
	Addr        string
	UserRpcConf zrpc.RpcClientConf
	MsgRpcConf  zrpc.RpcClientConf
}
