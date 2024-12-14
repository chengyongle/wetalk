package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"wetalk/apps/social/rpc/internal/svc"
	"wetalk/apps/social/rpc/social"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendListLogic {
	return &FriendListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *FriendListLogic) FriendList(in *social.FriendListReq) (*social.FriendListResp, error) {
	friendsList, err := l.svcCtx.FriendsModel.ListByUserid(l.ctx, in.UserId)
	if err != nil {
		logx.Errorf("list friend by uid err %v req %v ", err, in.UserId)
		return nil, err
	}

	var respList []*social.Friends
	copier.Copy(&respList, &friendsList)

	return &social.FriendListResp{
		List: respList,
	}, nil
}