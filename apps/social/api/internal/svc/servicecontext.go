package svc

import (
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
	"wetalk/apps/im/rpc/imclient"
	"wetalk/apps/social/api/internal/config"
	"wetalk/apps/social/rpc/socialclient"
	"wetalk/apps/user/rpc/userclient"
	"wetalk/pkg/interceptor"
	"wetalk/pkg/middleware"
)

type ServiceContext struct {
	Config                config.Config
	IdempotenceMiddleware rest.Middleware
	LimitMiddleware       rest.Middleware
	*redis.Redis
	socialclient.Social
	userclient.User
	imclient.Im
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:                c,
		Redis:                 redis.MustNewRedis(c.Redisx),
		IdempotenceMiddleware: middleware.NewIdempotenceMiddleware().Handler,
		LimitMiddleware:       middleware.NewLimitMiddleware(c.Redisx).TokenLimitHandler(1, 100),
		Social: socialclient.NewSocial(zrpc.MustNewClient(c.SocialRpc,
			zrpc.WithUnaryClientInterceptor(interceptor.DefaultIdempotentClient),
		)),
		User: userclient.NewUser(zrpc.MustNewClient(c.UserRpc)),
		Im:   imclient.NewIm(zrpc.MustNewClient(c.ImRpc)),
	}
}
