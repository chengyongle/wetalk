package logic

import (
	"context"
	"github.com/jinzhu/copier"

	"wetalk/apps/social/rpc/internal/svc"
	"wetalk/apps/social/rpc/social"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendPutInListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFriendPutInListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendPutInListLogic {
	return &FriendPutInListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 好友申请列表（别人对我的的申请）
func (l *FriendPutInListLogic) FriendPutInList(in *social.FriendPutInListReq) (*social.FriendPutInListResp, error) {
	friendReqList, err := l.svcCtx.FriendRequestsModel.ListNoHandler(l.ctx, in.UserId)
	if err != nil {
		logx.Errorf("list friend  req by uid err %v req %v ", err, in.UserId)
		return nil, err
	}

	var respList []*social.FriendRequests
	copier.Copy(&respList, &friendReqList)

	return &social.FriendPutInListResp{
		List: respList,
	}, nil
}
