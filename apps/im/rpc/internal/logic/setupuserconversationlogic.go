package logic

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"wetalk/apps/im/immodels"
	"wetalk/pkg/constants"
	"wetalk/pkg/wuid"

	"wetalk/apps/im/rpc/im"
	"wetalk/apps/im/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type SetUpUserConversationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSetUpUserConversationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SetUpUserConversationLogic {
	return &SetUpUserConversationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 建立会话: 群聊, 私聊
func (l *SetUpUserConversationLogic) SetUpUserConversation(in *im.SetUpUserConversationReq) (*im.SetUpUserConversationResp, error) {
	var res im.SetUpUserConversationResp
	switch constants.ChatType(in.ChatType) {
	case constants.SingleChatType:
		// 生成会话id
		conversationId := wuid.CombineId(in.SendId, in.RecvId)
		// 验证是否建立过会话
		conversationRes, err := l.svcCtx.ConversationModel.FindOne(l.ctx, conversationId)
		if err != nil {
			// 建立会话
			if err == immodels.ErrNotFound {
				err = l.svcCtx.ConversationModel.Insert(l.ctx, &immodels.Conversation{
					ConversationId: conversationId,
					ChatType:       constants.SingleChatType,
				})
				if err != nil {
					logx.Errorf("ConversationsModel.Insert err %v", err)
					return nil, err
				}
			} else {
				logx.Errorf("ConversationsModel.FindOne err %v, req %v", err, conversationId)
				return nil, err
			}
		} else if conversationRes != nil {
			return &res, nil
		}
		// 建立两者的会话
		err = l.setUpUserConversation(conversationId, in.SendId, in.RecvId, constants.SingleChatType, true)
		if err != nil {
			return nil, err
		}
		err = l.setUpUserConversation(conversationId, in.RecvId, in.SendId, constants.SingleChatType, false)
		if err != nil {
			return nil, err
		}
	case constants.GroupChatType:
		err := l.setUpUserConversation(in.RecvId, in.SendId, in.RecvId, constants.GroupChatType, true)
		if err != nil {
			return nil, err
		}
	}

	return &im.SetUpUserConversationResp{}, nil
}

// 新建会话记录
func (l *SetUpUserConversationLogic) setUpUserConversation(conversationId, userId, recvId string,
	chatType constants.ChatType, isShow bool) error {
	// 用户的会话列表
	conversations, err := l.svcCtx.ConversationsModel.FindByUserId(l.ctx, userId)
	if err != nil {
		if err == immodels.ErrNotFound {
			conversations = &immodels.Conversations{
				ID:               primitive.NewObjectID(),
				UserId:           userId,
				ConversationList: make(map[string]*immodels.Conversation),
			}
		} else {
			logx.Errorf("ConversationsModel.FindOne err %v, req %v", err, userId)
			return err
		}
	}
	// 更新会话记录
	if _, ok := conversations.ConversationList[conversationId]; ok {
		return nil
	}
	// 添加会话记录
	conversations.ConversationList[conversationId] = &immodels.Conversation{
		ConversationId: conversationId,
		ChatType:       chatType,
		IsShow:         isShow,
	}
	// 更新
	_, err = l.svcCtx.ConversationsModel.Update(l.ctx, conversations)
	if err != nil {
		logx.Errorf("ConversationsModel.Insert err %v, req %v", err, conversations)
		return err
	}
	return nil
}
