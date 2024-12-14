package group

import (
	"context"
	"github.com/jinzhu/copier"
	"wetalk/apps/social/rpc/socialclient"

	"wetalk/apps/social/api/internal/svc"
	"wetalk/apps/social/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupPutinListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 申请进群列表
func NewGroupPutinListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupPutinListLogic {
	return &GroupPutinListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupPutinListLogic) GroupPutinList(req *types.GroupPutinListReq) (resp *types.GroupPutinListResp, err error) {
	list, err := l.svcCtx.Social.GroupPutinList(l.ctx, &socialclient.GroupPutinListReq{
		GroupId: req.GroupId,
	})
	if err != nil {
		return nil, err
	}
	var respList []*types.GroupRequests
	copier.Copy(&respList, list.List)

	return &types.GroupPutinListResp{List: respList}, nil
}
