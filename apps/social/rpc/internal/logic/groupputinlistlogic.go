package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"wetalk/apps/social/rpc/internal/svc"
	"wetalk/apps/social/rpc/social"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupPutinListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGroupPutinListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupPutinListLogic {
	return &GroupPutinListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GroupPutinListLogic) GroupPutinList(in *social.GroupPutinListReq) (*social.GroupPutinListResp, error) {
	groupReqs, err := l.svcCtx.GroupRequestsModel.ListNoHandler(l.ctx, in.GroupId)
	if err != nil {
		logx.Errorf("list group req err %v req %v", err, in.GroupId)
		return nil, err
	}
	var respList []*social.GroupRequests
	copier.Copy(&respList, groupReqs)

	return &social.GroupPutinListResp{
		List: respList,
	}, nil
}
