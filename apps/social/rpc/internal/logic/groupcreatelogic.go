package logic

import (
	"context"
	"database/sql"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"
	"wetalk/apps/social/socialmodels"
	"wetalk/pkg/constants"
	"wetalk/pkg/wuid"

	"wetalk/apps/social/rpc/internal/svc"
	"wetalk/apps/social/rpc/social"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupCreateLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGroupCreateLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupCreateLogic {
	return &GroupCreateLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 群业务
func (l *GroupCreateLogic) GroupCreate(in *social.GroupCreateReq) (*social.GroupCreateResp, error) {
	groups := &socialmodels.Groups{
		Id:          wuid.GenUid(l.svcCtx.Config.Mysql.DataSource),
		Name:        in.Name,
		Icon:        in.Icon,
		CreatorUid:  in.CreatorUid,
		MemberNum:   1,
		IsVerify:    true, //加群是否要验证
		CreatedTime: time.Now(),
		UpdatedTime: time.Now(),
	}

	err := l.svcCtx.GroupsModel.Trans(l.ctx, func(ctx context.Context, session sqlx.Session) error {
		_, err := l.svcCtx.GroupsModel.Insert(l.ctx, session, groups)

		if err != nil {
			logx.Errorf("insert group err %v req %v", err, in)
			return err
		}

		_, err = l.svcCtx.GroupMembersModel.Insert(l.ctx, session, &socialmodels.GroupMembers{
			GroupId:      groups.Id,
			UserId:       in.CreatorUid,
			RoleLevel:    int64(constants.CreatorGroupRoleLevel),
			MemberStatus: 1,
			JoinTime: sql.NullTime{
				Time:  time.Now(),
				Valid: true,
			},
		})
		if err != nil {
			logx.Errorf("insert group member err %v req %v", err, in)
			return err
		}
		return nil
	})

	return &social.GroupCreateResp{Id: groups.Id}, err
}
