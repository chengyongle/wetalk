package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/conf"
	"wetalk/apps/im/ws/internal/config"
	"wetalk/apps/im/ws/internal/handler"
	"wetalk/apps/im/ws/internal/svc"
	"wetalk/apps/im/ws/websocket"
)

var configFile = flag.String("f", "etc/im.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	if err := c.SetUp(); err != nil {
		panic(err)
	}
	ctx := svc.NewServiceContext(c)
	srv := websocket.NewServer(c.ListenOn,
		//前置操作
		websocket.WithServerAuthentication(handler.NewJwtAuth(ctx)),
		//websocket.WithServerMaxConnectionIdle(5*time.Second),
		websocket.WithServerAck(websocket.NoAck),
	)
	defer srv.Stop()

	handler.RegisterHandlers(srv, ctx)

	fmt.Println("start websocket server at ", c.ListenOn)
	srv.Start()

}
