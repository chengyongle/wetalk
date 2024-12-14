package group

import (
	"context"
	"github.com/jinzhu/copier"
	"wetalk/apps/social/rpc/socialclient"

	"wetalk/apps/social/api/internal/svc"
	"wetalk/apps/social/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupExitLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 退出群
func NewGroupExitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupExitLogic {
	return &GroupExitLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupExitLogic) GroupExit(req *types.GroupExitReq) (*types.GroupExitResp, error) {
	// 退出群
	groupexitResp, err := l.svcCtx.Social.GroupExit(l.ctx, &socialclient.GroupExitReq{
		UserId:  req.UserId,
		GroupId: req.GroupId,
	})
	if err != nil {
		return nil, err
	}
	var resp types.GroupExitResp
	err = copier.Copy(&resp, groupexitResp)
	if err != nil {
		logx.Errorf("copier.Copy err %v", err)
		return nil, err
	}
	return &resp, err
}
