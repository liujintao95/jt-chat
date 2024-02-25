package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	JwtAuth JwtConf
	DB      DBConf
}

type JwtConf struct {
	AccessSecret string
	AccessExpire int64
}

type DBConf struct {
	DataSource string
}
