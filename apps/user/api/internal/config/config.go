package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	JwtAuth     JwtConf
	UserRpcConf zrpc.RpcClientConf
}

type JwtConf struct {
	AccessSecret string
	AccessExpire int64
}
