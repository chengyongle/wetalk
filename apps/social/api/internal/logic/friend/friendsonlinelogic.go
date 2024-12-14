package friend

import (
	"context"
	"wetalk/apps/social/rpc/social"
	"wetalk/pkg/constants"
	"wetalk/pkg/ctxdata"

	"wetalk/apps/social/api/internal/svc"
	"wetalk/apps/social/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FriendsOnlineLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 好友在线情况
func NewFriendsOnlineLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FriendsOnlineLogic {
	return &FriendsOnlineLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FriendsOnlineLogic) FriendsOnline(req *types.FriendsOnlineReq) (resp *types.FriendsOnlineResp, err error) {
	uid := ctxdata.GetUId(l.ctx)
	friendList, err := l.svcCtx.Social.FriendList(l.ctx, &social.FriendListReq{
		UserId: uid,
	})
	if err != nil || len(friendList.List) == 0 {
		return resp, err
	}
	//从缓存中查询在线的用户
	uids := make([]string, 0, len(friendList.List))
	for _, friend := range friendList.List {
		uids = append(uids, friend.FriendUid)
	}
	onlines, err := l.svcCtx.Redis.Hgetall(constants.REDIS_ONLINE_USER)
	if err != nil {
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
	return &types.FriendsOnlineResp{
		OnlineList: resOnlineList,
	}, nil
}
