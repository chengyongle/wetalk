package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"
	"wetalk/apps/social/rpc/internal/code"
	"wetalk/apps/social/socialmodels"
	"wetalk/pkg/constants"

	"wetalk/apps/social/rpc/internal/svc"
	"wetalk/apps/social/rpc/social"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendPutInHandleLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFriendPutInHandleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendPutInHandleLogic {
	return &FriendPutInHandleLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FriendPutInHandleLogic) FriendPutInHandle(in *social.FriendPutInHandleReq) (*social.FriendPutInHandleResp, error) {
	// 获取好友申请记录
	firendReq, err := l.svcCtx.FriendRequestsModel.FindOne(l.ctx, int64(in.FriendReqId))
	if err != nil {
		logx.Errorf("find friendsRequest by friendReqid err %v req %v ", err, in.FriendReqId)
		return nil, err
	}

	// 验证是否有处理
	switch constants.HandlerResult(firendReq.HandleResult.Int64) {
	case constants.PassHandlerResult:
		return nil, code.FriendReqBeforePass
	case constants.RefuseHandlerResult:
		return nil, code.FriendReqBeforeRefuse
	}

	firendReq.HandleResult.Int64 = int64(in.HandleResult)

	// 修改申请结果 -》 通过【建立两条好友关系记录】 -》 事务
	err = l.svcCtx.FriendRequestsModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		if err := l.svcCtx.FriendRequestsModel.Update(l.ctx, session, firendReq); err != nil {
			logx.Errorf("update friend request err %v, req %v", err, firendReq)
			return err
		}

		if constants.HandlerResult(in.HandleResult) != constants.PassHandlerResult {
			return nil
		}

		friends := []*socialmodels.Friends{
			{
				UserId:       firendReq.UserId,
				FriendUid:    firendReq.ReqUid,
				FriendStatus: 1,
				CreatedTime:  time.Now(),
				UpdatedTime:  time.Now(),
			}, {
				UserId:       firendReq.ReqUid,
				FriendUid:    firendReq.UserId,
				FriendStatus: 1,
				CreatedTime:  time.Now(),
				UpdatedTime:  time.Now(),
			},
		}

		_, err = l.svcCtx.FriendsModel.Inserts(l.ctx, session, friends...)
		if err != nil {
			logx.Errorf("friends inserts err %v, req %v", err, friends)
			return err
		}
		return nil
	})

	return &social.FriendPutInHandleResp{Success: true}, err
}
