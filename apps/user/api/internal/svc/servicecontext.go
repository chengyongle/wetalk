package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"wetalk/apps/user/api/internal/config"
	"wetalk/apps/user/rpc/userclient"
)

type ServiceContext struct {
	Config config.Config
	userclient.User
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		User:   userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
	}
}
