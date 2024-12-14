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

type GroupPutInHandleLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 申请进群处理
func NewGroupPutInHandleLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupPutInHandleLogic {
	return &GroupPutInHandleLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupPutInHandleLogic) GroupPutInHandle(req *types.GroupPutInHandleReq) (*types.GroupPutInHandleResp, error) {
	uid := ctxdata.GetUId(l.ctx)
	GPIHResp, err := l.svcCtx.Social.GroupPutInHandle(l.ctx, &socialclient.GroupPutInHandleReq{
		GroupReqId:   req.GroupReqId,
		GroupId:      req.GroupId,
		HandleUid:    uid,
		HandleResult: req.HandleResult,
	})
	if err != nil {
		return nil, err
	}
	var res *types.GroupPutInHandleResp
	if constants.HandlerResult(req.HandleResult) != constants.PassHandlerResult {
		return res, nil
	}

	if GPIHResp.GroupId == "" {
		return res, nil
	}

	_, err = l.svcCtx.Im.SetUpUserConversation(l.ctx, &imclient.SetUpUserConversationReq{
		SendId:   uid,
		RecvId:   GPIHResp.GroupId,
		ChatType: int32(constants.GroupChatType),
	})
	if err != nil {
		return nil, err
	}
	res.GroupId = GPIHResp.GroupId
	return res, nil
}
