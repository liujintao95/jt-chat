package config

import (
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	service.ServiceConf
	UserRpcConf zrpc.RpcClientConf
	MsgRpcConf  zrpc.RpcClientConf
}
