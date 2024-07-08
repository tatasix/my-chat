// Code generated by goctl. DO NOT EDIT.

package model

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
	rightsTimesFieldNames          = builder.RawFieldNames(&RightsTimes{})
	rightsTimesRows                = strings.Join(rightsTimesFieldNames, ",")
	rightsTimesRowsExpectAutoSet   = strings.Join(stringx.Remove(rightsTimesFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	rightsTimesRowsWithPlaceHolder = strings.Join(stringx.Remove(rightsTimesFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheRightsTimesIdPrefix = "cache:rightsTimes:id:"
)

type (
	rightsTimesModel interface {
		Insert(ctx context.Context, data *RightsTimes) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*RightsTimes, error)
		Update(ctx context.Context, data *RightsTimes) error
		Delete(ctx context.Context, id int64) error
	}

	defaultRightsTimesModel struct {
		sqlc.CachedConn
		table string
	}

	RightsTimes struct {
		Id        int64        `db:"id"`
		RightsId  int64        `db:"rights_id"`
		User      string       `db:"user"`       // 用户标识
		OpenKfId  string       `db:"open_kf_id"` // 客服标识
		PayType   int64        `db:"pay_type"`   // 付款方式：1 次卡;2 月卡; 3年卡
		Period    int64        `db:"period"`     // 计算周期：1 按天算次数;2 不按时间计算次数
		Start     sql.NullTime `db:"start"`      // 开始时间
		End       sql.NullTime `db:"end"`        // 结束时间
		Total     int64        `db:"total"`      // 总次数
		Remain    int64        `db:"remain"`     // 剩余次数
		Used      int64        `db:"used"`       // 使用次数
		CreatedAt time.Time    `db:"created_at"` // 创建时间
		UpdatedAt time.Time    `db:"updated_at"` // 更新时间
		IsDeleted int64        `db:"is_deleted"` // 是否删除，0否，1是
	}
)

func newRightsTimesModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultRightsTimesModel {
	return &defaultRightsTimesModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`rights_times`",
	}
}

func (m *defaultRightsTimesModel) Delete(ctx context.Context, id int64) error {
	rightsTimesIdKey := fmt.Sprintf("%s%v", cacheRightsTimesIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, rightsTimesIdKey)
	return err
}

func (m *defaultRightsTimesModel) FindOne(ctx context.Context, id int64) (*RightsTimes, error) {
	rightsTimesIdKey := fmt.Sprintf("%s%v", cacheRightsTimesIdPrefix, id)
	var resp RightsTimes
	err := m.QueryRowCtx(ctx, &resp, rightsTimesIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", rightsTimesRows, m.table)
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

func (m *defaultRightsTimesModel) Insert(ctx context.Context, data *RightsTimes) (sql.Result, error) {
	rightsTimesIdKey := fmt.Sprintf("%s%v", cacheRightsTimesIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)", m.table, rightsTimesRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.RightsId, data.User, data.OpenKfId, data.PayType, data.Period, data.Start, data.End, data.Total, data.Remain, data.Used, data.IsDeleted)
	}, rightsTimesIdKey)
	return ret, err
}

func (m *defaultRightsTimesModel) Update(ctx context.Context, data *RightsTimes) error {
	rightsTimesIdKey := fmt.Sprintf("%s%v", cacheRightsTimesIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, rightsTimesRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.RightsId, data.User, data.OpenKfId, data.PayType, data.Period, data.Start, data.End, data.Total, data.Remain, data.Used, data.IsDeleted, data.Id)
	}, rightsTimesIdKey)
	return err
}

func (m *defaultRightsTimesModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheRightsTimesIdPrefix, primary)
}

func (m *defaultRightsTimesModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", rightsTimesRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultRightsTimesModel) tableName() string {
	return m.table
}
