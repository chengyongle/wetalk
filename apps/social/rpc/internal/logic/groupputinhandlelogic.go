package logic

import (
	"context"
	"database/sql"
	"errors"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"
	"wetalk/apps/social/rpc/internal/code"
	"wetalk/apps/social/rpc/internal/svc"
	"wetalk/apps/social/rpc/social"
	"wetalk/apps/social/socialmodels"
	"wetalk/pkg/constants"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupPutInHandleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGroupPutInHandleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupPutInHandleLogic {
	return &GroupPutInHandleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GroupPutInHandleLogic) GroupPutInHandle(in *social.GroupPutInHandleReq) (*social.GroupPutInHandleResp, error) {
	groupReq, err := l.svcCtx.GroupRequestsModel.FindOne(l.ctx, int64(in.GroupReqId))
	if err != nil && err != sql.ErrNoRows {
		logx.Errorf("find group req err %v req %v", err, in.GroupReqId)
		return nil, err
	}
	if errors.Is(err, sql.ErrNoRows) {
		return nil, code.InvalidApplyId
	}
	switch constants.HandlerResult(groupReq.HandleResult.Int64) {
	case constants.PassHandlerResult:
		return nil, code.GroupReqBeforePass
	case constants.RefuseHandlerResult:
		return nil, code.GroupReqBeforeRefuse
	}

	groupReq.HandleResult = sql.NullInt64{
		Int64: int64(in.HandleResult),
		Valid: true,
	}

	err = l.svcCtx.GroupRequestsModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		if err := l.svcCtx.GroupRequestsModel.Update(l.ctx, session, groupReq); err != nil {
			logx.Errorf("update group req err %v req %v", err, groupReq)
			return err
		}

		if constants.HandlerResult(groupReq.HandleResult.Int64) != constants.PassHandlerResult {
			return nil
		}

		groupMember := &socialmodels.GroupMembers{
			GroupId:      groupReq.GroupId,
			UserId:       groupReq.ReqId,
			RoleLevel:    int64(constants.AtLargeGroupRoleLevel),
			MemberStatus: 1,
			JoinTime: sql.NullTime{
				Time:  time.Now(),
				Valid: true,
			},
			JoinSource: sql.NullInt64{
				Int64: 2,
				Valid: true,
			},
			OperatorUid: sql.NullString{
				String: in.HandleUid,
				Valid:  true,
			},
		}
		_, err = l.svcCtx.GroupMembersModel.Insert(l.ctx, session, groupMember)
		if err != nil {
			logx.Errorf("insert groupmember err %v req %v", err, groupMember)
			return err
		}
		err = l.svcCtx.GroupsModel.IncreaseGroupMember(l.ctx, session, groupMember.GroupId)
		if err != nil {
			logx.Errorf("Increase group member err %v req %v", err, in)
			return err
		}
		return nil
	})

	return &social.GroupPutInHandleResp{GroupId: groupReq.GroupId}, err
}
