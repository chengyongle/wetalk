package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"wetalk/apps/im/immodels"

	"wetalk/apps/im/rpc/im"
	"wetalk/apps/im/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetConversationsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetConversationsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetConversationsLogic {
	return &GetConversationsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取会话
func (l *GetConversationsLogic) GetConversations(in *im.GetConversationsReq) (*im.GetConversationsResp, error) {
	// 根据用户查询用户会话列表
	data, err := l.svcCtx.ConversationsModel.FindByUserId(l.ctx, in.UserId)
	if err != nil {
		if err == immodels.ErrNotFound {
			return &im.GetConversationsResp{}, nil
		}
		logx.Errorf("ConversationsModel.FindByUserId err %v, req %v", err, in.UserId)
		return nil, err
	}
	var res im.GetConversationsResp
	copier.Copy(&res, &data)
	// 根据会话列表，查询具体的会话
	ids := make([]string, 0, len(data.ConversationList))
	for _, conversation := range data.ConversationList {
		ids = append(ids, conversation.ConversationId)
	}
	conversations, err := l.svcCtx.ConversationModel.ListByConversationIds(l.ctx, ids)
	if err != nil {
		logx.Errorf("ConversationModel.ListByConversationIds err %v, req %v", err, ids)
		return nil, err
	}
	//res存的是用户（前端）看到的状态，conversations存的是数据库的状态
	// 计算是否存在未读消息
	for _, conversation := range conversations {
		if _, ok := res.ConversationList[conversation.ConversationId]; !ok {
			continue
		}
		// 用户读取的消息量
		total := res.ConversationList[conversation.ConversationId].Total
		if total < int32(conversation.Total) {
			// 有新的消息（离线消息）
			res.ConversationList[conversation.ConversationId].Total = int32(conversation.Total)
			// 未读数量
			res.ConversationList[conversation.ConversationId].ToRead = int32(conversation.Total) - total
			// 更改当前会话为显示状态
			res.ConversationList[conversation.ConversationId].IsShow = true
		}
	}

	return &res, nil
}
