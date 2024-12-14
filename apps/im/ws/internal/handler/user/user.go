package user

import (
	"wetalk/apps/im/ws/internal/svc"
	"wetalk/apps/im/ws/websocket"
)

// 在线的用户
func Online(svc *svc.ServiceContext) websocket.HandlerFunc {
	return func(srv *websocket.Server, conn *websocket.Conn, msg *websocket.Message) {
		uids := srv.GetUsers()
		u := srv.GetUsers(conn)
		err := srv.Send(websocket.NewMessage(u[0], uids), conn)
		srv.Info("err ", err)
	}
}
