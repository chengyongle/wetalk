package logic

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"
	"wetalk/apps/social/rpc/internal/code"
	"wetalk/apps/social/socialmodels"
	"wetalk/pkg/constants"

	"wetalk/apps/social/rpc/internal/svc"
	"wetalk/apps/social/rpc/social"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupPutinLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGroupPutinLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupPutinLogic {
	return &GroupPutinLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GroupPutinLogic) GroupPutin(in *social.GroupPutinReq) (*social.GroupPutinResp, error) {
	//  1. 普通用户申请 ： 如果群无验证直接进入
	//  2. 群成员邀请： 如果群无验证直接进入
	//  3. 群管理员/群创建者邀请：直接进入群

	var (
		inviteGroupMember *socialmodels.GroupMembers
		userGroupMember   *socialmodels.GroupMembers
		groupInfo         *socialmodels.Groups

		err error
	)
	//查找群信息
	groupInfo, err = l.svcCtx.GroupsModel.FindOne(l.ctx, in.GroupId)
	if err != nil {
		logx.Errorf("find group by groud id err %v, req %v", err, in.GroupId)
		return nil, err
	}
	if groupInfo == nil || groupInfo.Status.Int64 == 2 {
		return nil, code.GroupNotExist
	}
	//检查是否在群
	userGroupMember, err = l.svcCtx.GroupMembersModel.FindByGroudIdAndUserId(l.ctx, in.ReqId, in.GroupId)
	if err != nil && err != socialmodels.ErrNotFound {
		logx.Errorf("find group member by groud id and  req id err %v, req %v, %v", err, in.GroupId, in.ReqId)
		return nil, err
	}
	if userGroupMember != nil && userGroupMember.MemberStatus == 2 {
		return nil, code.CannotJoinAgain
	}
	if userGroupMember != nil {
		return nil, code.AlreadyJoin
	}
	//检查是否申请
	groupReq, err := l.svcCtx.GroupRequestsModel.FindByGroupIdAndReqId(l.ctx, in.GroupId, in.ReqId)
	if err != nil && err != socialmodels.ErrNotFound {
		logx.Errorf("find group req by groud id and user id err %v, req %v, %v", err, in.GroupId, in.ReqId)
		return nil, err
	}
	if groupReq != nil {
		return nil, code.AlreadyApplied
	}
	//群请求结构体
	groupReq = &socialmodels.GroupRequests{
		ReqId:   in.ReqId,
		GroupId: in.GroupId,
		ReqMsg: sql.NullString{
			String: in.ReqMsg,
			Valid:  true,
		},
		ReqTime: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		JoinSource: sql.NullInt64{
			Int64: int64(in.JoinSource),
			Valid: true,
		},
		InviterUserId: sql.NullString{
			String: in.InviterUid,
			Valid:  true,
		},
		HandleResult: sql.NullInt64{
			Int64: int64(constants.NoHandlerResult),
			Valid: true,
		},
	}
	createGroupMember := func() {
		if err != nil {
			return
		}
		err = l.createGroupMember(in)
	}
	// 检查是否要验证
	if !groupInfo.IsVerify {
		fmt.Println(groupInfo.IsVerify)
		// 不需要
		defer createGroupMember()

		groupReq.HandleResult = sql.NullInt64{
			Int64: int64(constants.PassHandlerResult),
			Valid: true,
		}

		return l.createGroupReq(groupReq, true)
	}
	// 验证进群方式
	if constants.GroupJoinSource(in.JoinSource) == constants.PutInGroupJoinSource {
		// 申请
		return l.createGroupReq(groupReq, false)
	}
	//邀请
	inviteGroupMember, err = l.svcCtx.GroupMembersModel.FindByGroudIdAndUserId(l.ctx, in.InviterUid, in.GroupId)
	if err != nil && err != socialmodels.ErrNotFound {
		logx.Errorf("find group member by groud id and user id err %v, req %v", in.InviterUid, in.GroupId)
		return nil, err
	}
	if inviteGroupMember == nil || inviteGroupMember.MemberStatus == 2 {
		return nil, code.InvalidInvitation
	}
	//检查邀请者
	if constants.GroupRoleLevel(inviteGroupMember.RoleLevel) == constants.CreatorGroupRoleLevel || constants.
		GroupRoleLevel(inviteGroupMember.RoleLevel) == constants.ManagerGroupRoleLevel {
		// 是管理者或创建者邀请
		defer createGroupMember()

		groupReq.HandleResult = sql.NullInt64{
			Int64: int64(constants.PassHandlerResult),
			Valid: true,
		}
		groupReq.HandleUserId = sql.NullString{
			String: in.InviterUid,
			Valid:  true,
		}
		return l.createGroupReq(groupReq, true)
	}
	return l.createGroupReq(groupReq, false)
}

func (l *GroupPutinLogic) createGroupReq(groupReq *socialmodels.GroupRequests, isPass bool) (*social.GroupPutinResp, error) {

	_, err := l.svcCtx.GroupRequestsModel.Insert(l.ctx, groupReq)
	if err != nil {
		logx.Errorf("insert group req err %v req %v", err, groupReq)
		return nil, err
	}

	if isPass {
		return &social.GroupPutinResp{GroupId: groupReq.GroupId}, nil
	}

	return &social.GroupPutinResp{GroupId: groupReq.GroupId}, nil
}

func (l *GroupPutinLogic) createGroupMember(in *social.GroupPutinReq) error {
	groupMember := &socialmodels.GroupMembers{
		GroupId:      in.GroupId,
		UserId:       in.ReqId,
		RoleLevel:    int64(constants.AtLargeGroupRoleLevel),
		MemberStatus: 1,
		OperatorUid: sql.NullString{
			String: in.InviterUid,
			Valid:  true,
		},
	}
	//事务
	err := l.svcCtx.GroupsModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		_, err := l.svcCtx.GroupMembersModel.Insert(l.ctx, session, groupMember)
		if err != nil {
			logx.Errorf("insert group member err %v req %v", err, groupMember)
			return err
		}
		err = l.svcCtx.GroupsModel.IncreaseGroupMember(l.ctx, session, groupMember.GroupId)
		if err != nil {
			logx.Errorf("Increase group member err %v req %v", err, in)
			return err
		}
		return nil
	})

	return err
}
