package handler

import (
	"wetalk/apps/im/ws/internal/handler/conversation"
	"wetalk/apps/im/ws/internal/handler/push"
	"wetalk/apps/im/ws/internal/handler/user"
	"wetalk/apps/im/ws/internal/svc"
	"wetalk/apps/im/ws/websocket"
)

func RegisterHandlers(srv *websocket.Server, svc *svc.ServiceContext) {
	srv.AddRoutes([]websocket.Route{
		{
			Method:  "user.online",
			Handler: user.Online(svc),
		},
		{
			Method:  "conversation.chat",
			Handler: conversation.Chat(svc),
		},
		{
			Method:  "conversation.markChat",
			Handler: conversation.MarkRead(svc),
		},
		{
			Method:  "push",
			Handler: push.Push(svc),
		},
	})
}
