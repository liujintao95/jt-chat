package main

import (
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
	"jt-chat/apps/chatserver/internal/config"
	"jt-chat/apps/chatserver/internal/svc"
)

var configFile = flag.String("f", "apps/chatserver/etc/chatserver.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	// log、prometheus、trace、metricsUrl
	if err := c.SetUp(); err != nil {
		panic(err)
	}

	svcContext := svc.NewServiceContext(c)
}
