package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"wetalk/apps/social/rpc/internal/code"
	"wetalk/apps/social/socialmodels"

	"wetalk/apps/social/rpc/internal/svc"
	"wetalk/apps/social/rpc/social"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupExitLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGroupExitLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupExitLogic {
	return &GroupExitLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 退出群
func (l *GroupExitLogic) GroupExit(in *social.GroupExitReq) (*social.GroupExitResp, error) {
	//检查是否在群
	userGroupMember, err := l.svcCtx.GroupMembersModel.FindByGroudIdAndUserId(l.ctx, in.UserId, in.GroupId)
	if err != nil && err != socialmodels.ErrNotFound {
		logx.Errorf("find group member by groud id and  req id err %v, req %v, %v", err, in.GroupId, in.UserId)
		return nil, err
	}
	if userGroupMember == nil || userGroupMember.MemberStatus == 2 {
		return nil, code.UserNotExistGroup
	}
	userGroupMember.MemberStatus = 2
	//事务
	err = l.svcCtx.GroupsModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		err := l.svcCtx.GroupMembersModel.Update(l.ctx, session, userGroupMember)
		if err != nil {
			logx.Errorf("insert group member err %v req %v", err, userGroupMember)
			return err
		}
		err = l.svcCtx.GroupsModel.DecreaseGroupMember(l.ctx, session, in.GroupId)
		if err != nil {
			logx.Errorf("Increase group member err %v req %v", err, in.GroupId)
			return err
		}
		return nil
	})
	if err != nil {
		logx.Errorf("GroupExit Trans err %v ", err)
		return nil, err
	}
	return &social.GroupExitResp{Success: true}, nil
}
