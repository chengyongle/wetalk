package logic

import (
	"context"
	"github.com/jinzhu/copier"
	"github.com/zeromicro/go-zero/core/mr"
	"wetalk/apps/social/rpc/internal/svc"
	"wetalk/apps/social/rpc/social"
	"wetalk/apps/social/socialmodels"

	"github.com/zeromicro/go-zero/core/logx"
)

type GroupListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGroupListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GroupListLogic {
	return &GroupListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GroupListLogic) GroupList(in *social.GroupListReq) (*social.GroupListResp, error) {
	userGroup, err := l.svcCtx.GroupMembersModel.ListByUserId(l.ctx, in.UserId)
	if err != nil {
		logx.Errorf("list group member err %v req %v", err, in.UserId)
		return nil, err
	}
	if len(userGroup) == 0 {
		return &social.GroupListResp{}, nil
	}

	ids := make([]string, 0, len(userGroup))
	for _, v := range userGroup {
		ids = append(ids, v.GroupId)
	}

	groups, err := l.groupByIds(l.ctx, ids)
	if err != nil {
		logx.Errorf("list group err %v req %v", err, ids)
		return nil, err
	}

	var respList []*social.Groups
	copier.Copy(&respList, &groups)

	return &social.GroupListResp{
		List: respList,
	}, nil
}
func (l *GroupListLogic) groupByIds(ctx context.Context, groupIds []string) ([]*socialmodels.Groups, error) {
	groups, err := mr.MapReduce[string, *socialmodels.Groups, []*socialmodels.Groups](func(source chan<- string) {
		//生成数据
		for _, gid := range groupIds {
			source <- gid
		}
	}, func(id string, writer mr.Writer[*socialmodels.Groups], cancel func(error)) {
		//处理数据
		p, err := l.svcCtx.GroupsModel.FindOne(ctx, id)
		if err != nil {
			cancel(err)
			return
		}
		writer.Write(p)
	}, func(pipe <-chan *socialmodels.Groups, writer mr.Writer[[]*socialmodels.Groups], cancel func(error)) {
		//聚合数据返回
		var groups []*socialmodels.Groups
		for group := range pipe {
			groups = append(groups, group)
		}
		writer.Write(groups)
	})
	if err != nil {
		return nil, err
	}

	return groups, nil
}
