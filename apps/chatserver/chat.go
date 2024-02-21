package main

import (
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"jt-chat/apps/chatserver/internal/config"
	"jt-chat/apps/chatserver/internal/handler"
)

var configFile = flag.String("f", "apps/chatserver/etc/chatserver.yaml", "the config file")

func main() {
	var c config.Config
	conf.MustLoad(*configFile, &c)
	logx.MustSetup(c.Log)
	server := handler.NewSocketServer()
	server.Start()
	defer server.Stop()
}
