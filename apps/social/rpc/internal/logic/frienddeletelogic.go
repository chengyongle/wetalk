package logic

import (
	"context"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"
	"wetalk/apps/social/rpc/internal/code"
	"wetalk/apps/social/rpc/internal/svc"
	"wetalk/apps/social/rpc/social"
	"wetalk/apps/social/socialmodels"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendDeleteLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFriendDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendDeleteLogic {
	return &FriendDeleteLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除好友
func (l *FriendDeleteLogic) FriendDelete(in *social.FriendDeleteReq) (*social.FriendDeleteResp, error) {
	// 申请人是否与目标是好友关系
	friends1, err := l.svcCtx.FriendsModel.FindByUidAndFid(l.ctx, in.UserId, in.FriendUid)
	if err != nil && err != socialmodels.ErrNotFound {
		logx.Errorf("find friendsRequest by uid and fid err %v req %v ", err, in)
		return nil, err
	}
	if err == socialmodels.ErrNotFound {
		return nil, code.NotFriend
	}
	friends2, err := l.svcCtx.FriendsModel.FindByUidAndFid(l.ctx, in.FriendUid, in.UserId)
	if err != nil && err != socialmodels.ErrNotFound {
		logx.Errorf("find friendsRequest by fid and uid err %v req %v ", err, in)
		return nil, err
	}
	friends1.FriendStatus = 2
	friends1.UpdatedTime = time.Now()
	friends2.FriendStatus = 2
	friends2.UpdatedTime = time.Now()
	// 事务 修改好友状态
	err = l.svcCtx.FriendsModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		if err := l.svcCtx.FriendsModel.Update(l.ctx, session, friends1); err != nil {
			logx.Errorf("update friend1 err: %v", err)
			return err
		}
		if err := l.svcCtx.FriendsModel.Update(l.ctx, session, friends2); err != nil {
			logx.Errorf("update friend2 err: %v", err)
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &social.FriendDeleteResp{Success: true}, nil
}
