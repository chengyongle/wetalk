package group

import (
	"context"
	"wetalk/apps/social/api/internal/svc"
	"wetalk/apps/social/api/internal/types"
	"wetalk/apps/social/rpc/socialclient"
	"wetalk/pkg/constants"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupUserOnlineLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 群在线用户
func NewGroupUserOnlineLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupUserOnlineLogic {
	return &GroupUserOnlineLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupUserOnlineLogic) GroupUserOnline(req *types.GroupUserOnlineReq) (resp *types.GroupUserOnlineResp, err error) {
	groupUsers, err := l.svcCtx.Social.GroupUsers(l.ctx, &socialclient.GroupUsersReq{
		GroupId: req.GroupId,
	})

	//从缓存中查询在线的用户
	uids := make([]string, 0, len(groupUsers.List))
	for _, v := range groupUsers.List {
		uids = append(uids, v.UserId)
	}

	onlines, err := l.svcCtx.Redis.Hgetall(constants.REDIS_ONLINE_USER)
	if err != nil {
		logx.Errorf("redis hgetall %v err %v:", constants.REDIS_ONLINE_USER, err)
		return nil, err
	}
	resOnlineList := make(map[string]bool, len(uids))
	for _, s := range uids {
		if _, ok := onlines[s]; ok {
			resOnlineList[s] = true
		} else {
			resOnlineList[s] = false
		}
	}
	return &types.GroupUserOnlineResp{
		OnlineList: resOnlineList,
	}, nil
}
