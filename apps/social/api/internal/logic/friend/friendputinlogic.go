package friend

import (
	"context"
	"github.com/jinzhu/copier"
	"wetalk/apps/social/rpc/socialclient"
	"wetalk/pkg/ctxdata"

	"wetalk/apps/social/api/internal/svc"
	"wetalk/apps/social/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendPutInLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 好友申请
func NewFriendPutInLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendPutInLogic {
	return &FriendPutInLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FriendPutInLogic) FriendPutIn(req *types.FriendPutInReq) (*types.FriendPutInResp, error) {
	uid := ctxdata.GetUId(l.ctx)

	PutInResp, err := l.svcCtx.Social.FriendPutIn(l.ctx, &socialclient.FriendPutInReq{
		UserId: uid,
		ReqUid: req.UserId,
		ReqMsg: req.ReqMsg,
	})
	var resp types.FriendPutInResp
	err = copier.Copy(&resp, PutInResp)
	if err != nil {
		logx.Errorf("copier.Copy err %v", err)
		return nil, err
	}
	return &resp, err
}
