package model

import (
	"chat/common/util"
	"context"
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"
)

var _ ResourceUsageModel = (*customResourceUsageModel)(nil)

type (
	// ResourceUsageModel is an interface to be customized, add more methods here,
	// and implement the added methods in customResourceUsageModel.
	ResourceUsageModel interface {
		resourceUsageModel
		FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*ResourceUsage, error)
		RowBuilder() squirrel.SelectBuilder
		FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*ResourceUsage, error)
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error)
		CountBuilder(field string) squirrel.SelectBuilder
		BuildFiled(old, new *ResourceUsage)
		UpdateTimesAndToken(ctx context.Context, id, times, token int64) (err error)
	}

	customResourceUsageModel struct {
		*defaultResourceUsageModel
	}
)

// NewResourceUsageModel returns a model for the database table.
func NewResourceUsageModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ResourceUsageModel {
	return &customResourceUsageModel{
		defaultResourceUsageModel: newResourceUsageModel(conn, c, opts...),
	}
}

func (m *defaultResourceUsageModel) FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*ResourceUsage, error) {

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp ResourceUsage
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
func (m *defaultResourceUsageModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(resourceUsageRows).From(m.table)
}

func (m *defaultResourceUsageModel) FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*ResourceUsage, error) {

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*ResourceUsage
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultResourceUsageModel) FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error) {

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

func (m *defaultResourceUsageModel) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("COUNT(" + field + ")").From(m.table)
}

func (m *defaultResourceUsageModel) BuildFiled(old, new *ResourceUsage) {
	if new == nil {
		return
	}
	if old == nil && new.Id == 0 {
		new.Id = util.GenerateSnowflakeInt64()
		return
	}
	new.Id = old.Id
	//还需要处理

	new.UpdatedAt = time.Now()

	return
}

func (m *defaultResourceUsageModel) UpdateTimesAndToken(ctx context.Context, id, times, token int64) (err error) {

	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set `times`= `times` + ? ,`token`= `token` + ? where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, times, token, id)
	})

	return
}
