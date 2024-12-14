package logic

import (
	"context"
	"wetalk/apps/im/immodels"
	"wetalk/pkg/constants"

	"wetalk/apps/im/rpc/im"
	"wetalk/apps/im/rpc/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateGroupConversationLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateGroupConversationLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateGroupConversationLogic {
	return &CreateGroupConversationLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *CreateGroupConversationLogic) CreateGroupConversation(in *im.CreateGroupConversationReq) (res *im.CreateGroupConversationResp, err error) {
	_, err = l.svcCtx.ConversationModel.FindOne(l.ctx, in.GroupId)
	if err == nil {
		//如果存在记录，说明无需创建
		return res, nil
	}
	if err != immodels.ErrNotFound {
		logx.Errorf("ConversationModel.FindOne err %v, req %v", err, in.GroupId)
		return nil, err
	}
	err = l.svcCtx.ConversationModel.Insert(l.ctx, &immodels.Conversation{
		ConversationId: in.GroupId,
		ChatType:       constants.GroupChatType,
	})
	if err != nil {
		logx.Errorf("ConversationModel.Insert err %v, req %v", err, in.GroupId)
		return nil, err
	}
	_, err = NewSetUpUserConversationLogic(l.ctx, l.svcCtx).SetUpUserConversation(&im.SetUpUserConversationReq{
		SendId:   in.CreateId,
		RecvId:   in.GroupId,
		ChatType: int32(constants.GroupChatType),
	})
	return res, err
}
