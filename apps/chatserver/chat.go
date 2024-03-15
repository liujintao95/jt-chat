package main

import (
	"context"
	"flag"
	"github.com/zeromicro/go-zero/core/conf"
	"jt-chat/apps/chatserver/internal/config"
	"jt-chat/apps/chatserver/internal/service"
	"jt-chat/apps/chatserver/internal/svc"
	"net/http"
	"time"
)

var configFile = flag.String("f", "apps/chatserver/etc/chatserver.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	if err := c.SetUp(); err != nil {
		panic(err)
	}

	svcContext := svc.NewServiceContext(c)
	ctx := context.Background()

	hub := service.NewSocketHub(ctx, svcContext)
	go hub.Start()
	defer hub.Stop()

	http.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		service.ServeWs(r.Context(), hub, w, r)
	})
	server := &http.Server{
		Addr:              c.Addr,
		ReadHeaderTimeout: 3 * time.Second,
	}
	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
