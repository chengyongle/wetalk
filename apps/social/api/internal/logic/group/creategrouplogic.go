package group

import (
	"context"
	"wetalk/apps/im/rpc/imclient"
	"wetalk/apps/social/rpc/socialclient"
	"wetalk/pkg/ctxdata"

	"wetalk/apps/social/api/internal/svc"
	"wetalk/apps/social/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateGroupLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创群
func NewCreateGroupLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateGroupLogic {
	return &CreateGroupLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateGroupLogic) CreateGroup(req *types.GroupCreateReq) (*types.GroupCreateResp, error) {
	uid := ctxdata.GetUId(l.ctx)

	// 创建群
	GroupCreateResp, err := l.svcCtx.Social.GroupCreate(l.ctx, &socialclient.GroupCreateReq{
		Name:       req.Name,
		Icon:       req.Icon,
		CreatorUid: uid,
	})
	if err != nil {
		return nil, err
	}
	if GroupCreateResp.Id == "" {
		return nil, err
	}
	// 建立会话
	_, err = l.svcCtx.Im.CreateGroupConversation(l.ctx, &imclient.CreateGroupConversationReq{
		GroupId:  GroupCreateResp.Id,
		CreateId: uid,
	})
	if err != nil {
		return nil, err
	}
	return &types.GroupCreateResp{Id: GroupCreateResp.Id}, err
}
