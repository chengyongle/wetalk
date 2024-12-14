package group

import (
	"context"
	"wetalk/apps/im/rpc/imclient"
	"wetalk/apps/social/rpc/socialclient"
	"wetalk/pkg/constants"
	"wetalk/pkg/ctxdata"

	"wetalk/apps/social/api/internal/svc"
	"wetalk/apps/social/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupPutinLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 申请进群
func NewGroupPutinLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupPutinLogic {
	return &GroupPutinLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupPutinLogic) GroupPutin(req *types.GroupPutinReq) (*types.GroupPutinResp, error) {
	uid := ctxdata.GetUId(l.ctx)

	PutinResp, err := l.svcCtx.Social.GroupPutin(l.ctx, &socialclient.GroupPutinReq{
		GroupId:    req.GroupId,
		ReqId:      uid,
		ReqMsg:     req.ReqMsg,
		JoinSource: req.JoinSource,
		InviterUid: req.InviterUid,
	})
	if err != nil {
		return nil, err
	}
	if PutinResp.GroupId == "" {
		return nil, err
	}
	// 建立会话
	_, err = l.svcCtx.Im.SetUpUserConversation(l.ctx, &imclient.SetUpUserConversationReq{
		SendId:   uid,
		RecvId:   PutinResp.GroupId,
		ChatType: int32(constants.GroupChatType),
	})
	return &types.GroupPutinResp{GroupId: PutinResp.GroupId}, err
}
