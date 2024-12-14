package svc

import (
	"github.com/zeromicro/go-zero/zrpc"
	"wetalk/apps/im/rpc/imclient"
	"wetalk/apps/social/api/internal/config"
	"wetalk/apps/social/rpc/socialclient"
	"wetalk/apps/user/rpc/userclient"
)

type ServiceContext struct {
	Config config.Config

	socialclient.Social
	userclient.User
	imclient.Im
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Social: socialclient.NewSocial(zrpc.MustNewClient(c.SocialRpc)),
		User:   userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		Im:     imclient.NewIm(zrpc.MustNewClient(c.ImRpc)),
	}
}
