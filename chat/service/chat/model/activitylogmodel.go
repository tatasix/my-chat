package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"
)

var _ ActivityLogModel = (*customActivityLogModel)(nil)

type (
	// ActivityLogModel is an interface to be customized, add more methods here,
	// and implement the added methods in customActivityLogModel.
	ActivityLogModel interface {
		activityLogModel
		FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*ActivityLog, error)
		RowBuilder() squirrel.SelectBuilder
		FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*ActivityLog, error)
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error)
		CountBuilder(field string) squirrel.SelectBuilder
		BuildFiled(old, new *ActivityLog)
		UpdateUser(ctx context.Context, newUser, oldUser string) error
	}

	customActivityLogModel struct {
		*defaultActivityLogModel
	}
)

// NewActivityLogModel returns a model for the database table.
func NewActivityLogModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ActivityLogModel {
	return &customActivityLogModel{
		defaultActivityLogModel: newActivityLogModel(conn, c, opts...),
	}
}

func (m *defaultActivityLogModel) FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*ActivityLog, error) {

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp ActivityLog
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &resp, nil
}

// export logic
func (m *defaultActivityLogModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(activityLogRows).From(m.table)
}

func (m *defaultActivityLogModel) FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*ActivityLog, error) {

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*ActivityLog
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultActivityLogModel) FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error) {

	query, values, err := countBuilder.ToSql()
	if err != nil {
		return 0, err
	}

	var resp int64
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return 0, nil
		}
		return 0, err
	}
	return resp, nil
}

func (m *defaultActivityLogModel) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("COUNT(" + field + ")").From(m.table)
}

func (m *defaultActivityLogModel) BuildFiled(old, new *ActivityLog) {
	if new == nil {
		return
	}
	new.UpdatedAt = time.Now()
	new.CreatedAt = time.Now()

	return
}

func (m *defaultActivityLogModel) UpdateUser(ctx context.Context, newUser, oldUser string) (err error) {
	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set `old_user`= `user` ,`user`= ? where `user` = ?", m.table)
		return conn.ExecCtx(ctx, query, newUser, oldUser)
	})

	return
}
