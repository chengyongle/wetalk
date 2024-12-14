package logic

import (
	"context"
	"database/sql"
	"time"
	"wetalk/apps/social/rpc/internal/code"
	"wetalk/apps/social/socialmodels"
	"wetalk/pkg/constants"

	"wetalk/apps/social/rpc/internal/svc"
	"wetalk/apps/social/rpc/social"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendPutInLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFriendPutInLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendPutInLogic {
	return &FriendPutInLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FriendPutInLogic) FriendPutIn(in *social.FriendPutInReq) (*social.FriendPutInResp, error) {
	// 申请人是否与目标是好友关系
	friends, err := l.svcCtx.FriendsModel.FindByUidAndFid(l.ctx, in.UserId, in.ReqUid)
	if err != nil && err != socialmodels.ErrNotFound {
		logx.Errorf("find friendsRequest by rid and uid err %v req %v ", err, in)
		return nil, err
	}
	if friends != nil && friends.FriendStatus == 2 {
		return nil, code.CannotReapply
	}
	if friends != nil {
		return nil, code.AlreadyFriend
	}
	// 是否已经有过申请
	friendReqs, err := l.svcCtx.FriendRequestsModel.FindByReqUidAndUserId(l.ctx, in.ReqUid, in.UserId)
	if err != nil && err != socialmodels.ErrNotFound {
		logx.Errorf("find friends by uid and fid err %v req %v ", err, in)
		return nil, err
	}
	if friendReqs != nil {
		return nil, code.AlreadyApplied
	}

	// 创建申请记录
	if in.ReqTime == 0 {
		in.ReqTime = time.Now().Unix()
	}
	_, err = l.svcCtx.FriendRequestsModel.Insert(l.ctx, &socialmodels.FriendRequests{
		UserId: in.UserId,
		ReqUid: in.ReqUid,
		ReqMsg: sql.NullString{
			Valid:  true,
			String: in.ReqMsg,
		},
		ReqTime: time.Unix(in.ReqTime, 0),
		HandleResult: sql.NullInt64{
			Int64: int64(constants.NoHandlerResult),
			Valid: true,
		},
	})

	if err != nil {
		logx.Errorf("insert friendRequest err %v req %v ", err, in)
		return nil, err
	}

	return &social.FriendPutInResp{Success: true}, nil
}
