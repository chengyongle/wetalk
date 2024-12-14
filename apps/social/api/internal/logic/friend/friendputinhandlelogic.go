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

type FriendPutInHandleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 好友申请处理
func NewFriendPutInHandleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendPutInHandleLogic {
	return &FriendPutInHandleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FriendPutInHandleLogic) FriendPutInHandle(req *types.FriendPutInHandleReq) (*types.FriendPutInHandleResp, error) {
	puthandleResp, err := l.svcCtx.Social.FriendPutInHandle(l.ctx, &socialclient.FriendPutInHandleReq{
		FriendReqId:  req.FriendReqId,
		UserId:       ctxdata.GetUId(l.ctx),
		HandleResult: req.HandleResult,
	})
	var resp types.FriendPutInHandleResp
	err = copier.Copy(&resp, puthandleResp)
	if err != nil {
		logx.Errorf("copier.Copy err %v", err)
		return nil, err
	}
	return &resp, err
}
