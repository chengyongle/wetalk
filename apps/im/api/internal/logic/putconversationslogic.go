package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"wetalk/apps/im/rpc/imclient"
	"wetalk/pkg/ctxdata"

	"wetalk/apps/im/api/internal/svc"
	"wetalk/apps/im/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutConversationsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新会话
func NewPutConversationsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutConversationsLogic {
	return &PutConversationsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PutConversationsLogic) PutConversations(req *types.PutConversationsReq) (resp *types.PutConversationsResp, err error) {
	uid := ctxdata.GetUId(l.ctx)
	var conversationList map[string]*imclient.Conversation
	copier.Copy(&conversationList, req.ConversationList)
	_, err = l.svcCtx.Im.PutConversations(l.ctx, &imclient.PutConversationsReq{
		UserId:           uid,
		ConversationList: conversationList,
	})
	if err != nil {
		return nil, err
	}
	return
}
