package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"wetalk/apps/user/rpc/user"

	"wetalk/apps/user/api/internal/svc"
	"wetalk/apps/user/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindUserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFindUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindUserLogic {
	return &FindUserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FindUserLogic) FindUser(req *types.FindUserRequest) (resp *types.FindUserResponse, err error) {
	finduserResp, err := l.svcCtx.User.FindUser(l.ctx, &user.FindUserReq{
		Name:  req.Name,
		Phone: req.Phone,
		Ids:   req.Ids,
	})
	if err != nil {
		return nil, err
	}

	var res types.FindUserResponse
	copier.Copy(&res, finduserResp)

	return &res, nil
}
