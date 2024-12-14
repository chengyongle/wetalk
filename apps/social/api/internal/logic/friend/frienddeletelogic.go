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

type FriendDeleteLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除好友
func NewFriendDeleteLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendDeleteLogic {
	return &FriendDeleteLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FriendDeleteLogic) FriendDelete(req *types.FriendDeleteReq) (*types.FriendDeleteResp, error) {
	uid := ctxdata.GetUId(l.ctx)

	deleteResp, err := l.svcCtx.Social.FriendDelete(l.ctx, &socialclient.FriendDeleteReq{
		UserId:    uid,
		FriendUid: req.FriendUid,
	})
	if err != nil {
		return nil, err
	}
	var resp types.FriendDeleteResp
	err = copier.Copy(&resp, deleteResp)
	if err != nil {
		logx.Errorf("copier.Copy err %v", err)
		return nil, err
	}
	return &resp, err
}
