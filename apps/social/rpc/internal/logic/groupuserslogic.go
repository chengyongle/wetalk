package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"wetalk/apps/social/rpc/internal/svc"
	"wetalk/apps/social/rpc/social"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupUsersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGroupUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupUsersLogic {
	return &GroupUsersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GroupUsersLogic) GroupUsers(in *social.GroupUsersReq) (*social.GroupUsersResp, error) {
	groupMembers, err := l.svcCtx.GroupMembersModel.ListByGroupId(l.ctx, in.GroupId)
	if err != nil {
		logx.Errorf("list group member err %v req %v", err, in.GroupId)
		return nil, err
	}

	var respList []*social.GroupMembers
	copier.Copy(&respList, &groupMembers)

	return &social.GroupUsersResp{
		List: respList,
	}, nil
}
