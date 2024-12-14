package logic

import (
	"context"
	"wetalk/apps/im/immodels"
	"wetalk/pkg/constants"

	"wetalk/apps/im/rpc/im"
	"wetalk/apps/im/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type PutConversationsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPutConversationsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutConversationsLogic {
	return &PutConversationsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新会话
func (l *PutConversationsLogic) PutConversations(in *im.PutConversationsReq) (*im.PutConversationsResp, error) {
	//查询会话列表
	data, err := l.svcCtx.ConversationsModel.FindByUserId(l.ctx, in.UserId)
	if err != nil {
		logx.Errorf("ConversationsModel.FindByUserId err %v, req %v", err, in.UserId)
		return nil, err
	}

	if data.ConversationList == nil {
		data.ConversationList = make(map[string]*immodels.Conversation)
	}
	//逐一更新会话列表
	for s, conversation := range in.ConversationList {
		var oldTotal int
		if data.ConversationList[s] != nil {
			oldTotal = data.ConversationList[s].Total
		}
		data.ConversationList[s] = &immodels.Conversation{
			ConversationId: conversation.ConversationId,
			ChatType:       constants.ChatType(conversation.ChatType),
			IsShow:         conversation.IsShow,
			Total:          int(conversation.Read) + oldTotal,
			Seq:            conversation.Seq,
		}
	}
	_, err = l.svcCtx.ConversationsModel.Update(l.ctx, data)
	if err != nil {
		logx.Errorf("ConversationsModel.Update err %v, req %v", err, data)
		return nil, err
	}

	return &im.PutConversationsResp{}, nil
}
