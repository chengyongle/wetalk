package logic

import (
	"context"
	"wetalk/apps/user/rpc/internal/code"

	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/logx"
	"wetalk/apps/user/models"
	"wetalk/apps/user/rpc/internal/svc"
	"wetalk/apps/user/rpc/user"
)

type GetUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserInfoLogic {
	return &GetUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetUserInfoLogic) GetUserInfo(in *user.GetUserInfoReq) (*user.GetUserInfoResp, error) {
	userEntiy, err := l.svcCtx.UsersModel.FindOne(l.ctx, in.Id)
	if err != nil {
		if err == models.ErrNotFound {
			return nil, code.UserNotFound
		}
		logx.Errorf("GetUserInfo  err %v ", err)
		return nil, err
	}
	var resp user.UserEntity
	//复制响应结构体
	copier.Copy(&resp, userEntiy)

	return &user.GetUserInfoResp{
		User: &resp,
	}, nil
}
