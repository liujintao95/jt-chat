package config

import (
	"github.com/zeromicro/go-zero/rest"
)

type Config struct {
	rest.RestConf
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
