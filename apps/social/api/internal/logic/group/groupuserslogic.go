package group

import (
	"context"
	"wetalk/apps/social/rpc/socialclient"
	"wetalk/apps/user/rpc/userclient"

	"wetalk/apps/social/api/internal/svc"
	"wetalk/apps/social/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupUsersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 成员列表
func NewGroupUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupUsersLogic {
	return &GroupUsersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GroupUsersLogic) GroupUsers(req *types.GroupUsersReq) (resp *types.GroupUsersResp, err error) {
	groupUsers, err := l.svcCtx.Social.GroupUsers(l.ctx, &socialclient.GroupUsersReq{
		GroupId: req.GroupId,
	})

	// 还需要获取用户的信息
	uids := make([]string, 0, len(groupUsers.List))
	for _, v := range groupUsers.List {
		uids = append(uids, v.UserId)
	}

	// 获取用户信息
	userList, err := l.svcCtx.User.FindUser(l.ctx, &userclient.FindUserReq{Ids: uids})
	if err != nil {
		return nil, err
	}

	userRecords := make(map[string]*userclient.UserEntity, len(userList.User))
	for i, _ := range userList.User {
		userRecords[userList.User[i].Id] = userList.User[i]
	}

	respList := make([]*types.GroupMembers, 0, len(groupUsers.List))
	for _, v := range groupUsers.List {

		member := &types.GroupMembers{
			Id:        v.Id,
			GroupId:   v.GroupId,
			UserId:    v.UserId,
			RoleLevel: v.RoleLevel,
		}
		if u, ok := userRecords[v.UserId]; ok {
			member.Nickname = u.Nickname
			member.UserAvatarUrl = u.Avatar
		}
		respList = append(respList, member)
	}

	return &types.GroupUsersResp{List: respList}, err
}
