// Code generated by goctl. DO NOT EDIT.
// versions:
//  goctl version: 1.7.2

package socialmodels

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	groupsFieldNames          = builder.RawFieldNames(&Groups{})
	groupsRows                = strings.Join(groupsFieldNames, ",")
	groupsRowsExpectAutoSet   = strings.Join(stringx.Remove(groupsFieldNames, "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	groupsRowsWithPlaceHolder = strings.Join(stringx.Remove(groupsFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheGroupsIdPrefix = "cache:groups:id:"
)

type (
	groupsModel interface {
		Trans(ctx context.Context, fn func(context.Context, sqlx.Session) error) error
		Insert(ctx context.Context, session sqlx.Session, data *Groups) (sql.Result, error)
		FindOne(ctx context.Context, id string) (*Groups, error)
		ListByGroupIds(ctx context.Context, ids []string) ([]*Groups, error)
		IncreaseGroupMember(ctx context.Context, session sqlx.Session, groupId string) error
		DecreaseGroupMember(ctx context.Context, session sqlx.Session, groupId string) error
		Update(ctx context.Context, data *Groups) error
		Delete(ctx context.Context, id string) error
	}

	defaultGroupsModel struct {
		sqlc.CachedConn
		table string
	}

	Groups struct {
		Id              string         `db:"id"`
		Name            string         `db:"name"`
		Icon            string         `db:"icon"`
		Status          sql.NullInt64  `db:"status"`
		CreatorUid      string         `db:"creator_uid"`
		GroupType       int64          `db:"group_type"`
		MemberNum       int64          `db:"member_num"`
		IsVerify        bool           `db:"is_verify"`
		Notification    sql.NullString `db:"notification"`
		NotificationUid sql.NullString `db:"notification_uid"`
		CreatedTime     time.Time      `db:"created_time"`
		UpdatedTime     time.Time      `db:"updated_time"`
	}
)

func newGroupsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultGroupsModel {
	return &defaultGroupsModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`groups`",
	}
}

func (m *defaultGroupsModel) Trans(ctx context.Context, fn func(context.Context, sqlx.Session) error) error {
	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})
}

func (m *defaultGroupsModel) Delete(ctx context.Context, id string) error {
	groupsIdKey := fmt.Sprintf("%s%v", cacheGroupsIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, groupsIdKey)
	return err
}

func (m *defaultGroupsModel) FindOne(ctx context.Context, id string) (*Groups, error) {
	groupsIdKey := fmt.Sprintf("%s%v", cacheGroupsIdPrefix, id)
	var resp Groups
	err := m.QueryRowCtx(ctx, &resp, groupsIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", groupsRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultGroupsModel) Insert(ctx context.Context, session sqlx.Session, data *Groups) (sql.Result, error) {
	groupsIdKey := fmt.Sprintf("%s%v", cacheGroupsIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, groupsRowsExpectAutoSet)
		return session.ExecCtx(ctx, query, data.Id, data.Name, data.Icon, data.Status, data.CreatorUid, data.GroupType, data.MemberNum, data.IsVerify, data.Notification, data.NotificationUid, data.CreatedTime, data.UpdatedTime)
	}, groupsIdKey)
	return ret, err
}

func (m *defaultGroupsModel) ListByGroupIds(ctx context.Context, ids []string) ([]*Groups, error) {
	query := fmt.Sprintf("select %s from %s where `id` in (?)", groupsRows, m.table)

	var resp []*Groups

	idStr := strings.Join(ids, "','")
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, idStr)

	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultGroupsModel) Update(ctx context.Context, data *Groups) error {
	groupsIdKey := fmt.Sprintf("%s%v", cacheGroupsIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, groupsRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Name, data.Icon, data.Status, data.CreatorUid, data.GroupType, data.MemberNum, data.IsVerify, data.Notification, data.NotificationUid, data.CreatedTime, data.UpdatedTime, data.Id)
	}, groupsIdKey)
	return err
}

func (m *defaultGroupsModel) IncreaseGroupMember(ctx context.Context, session sqlx.Session, groupId string) error {
	groupsIdKey := fmt.Sprintf("%s%v", cacheGroupsIdPrefix, groupId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set member_num=member_num+1 where `id` = ?", m.table)
		return session.ExecCtx(ctx, query, groupId)
	}, groupsIdKey)
	return err
}

func (m *defaultGroupsModel) DecreaseGroupMember(ctx context.Context, session sqlx.Session, groupId string) error {
	groupsIdKey := fmt.Sprintf("%s%v", cacheGroupsIdPrefix, groupId)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set member_num=member_num-1 where `id` = ?", m.table)
		return session.ExecCtx(ctx, query, groupId)
	}, groupsIdKey)
	return err
}

func (m *defaultGroupsModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheGroupsIdPrefix, primary)
}

func (m *defaultGroupsModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", groupsRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultGroupsModel) tableName() string {
	return m.table
}