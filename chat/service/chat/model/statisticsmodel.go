package model

import (
	"chat/common/util"
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"
)

var _ StatisticsModel = (*customStatisticsModel)(nil)

type (
	// StatisticsModel is an interface to be customized, add more methods here,
	// and implement the added methods in customStatisticsModel.
	StatisticsModel interface {
		statisticsModel
		FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*Statistics, error)
		RowBuilder() squirrel.SelectBuilder
		FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*Statistics, error)
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error)
		CountBuilder(field string) squirrel.SelectBuilder
		BuildFiled(old, new *Statistics)
	}

	customStatisticsModel struct {
		*defaultStatisticsModel
	}
)

// NewStatisticsModel returns a model for the database table.
func NewStatisticsModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) StatisticsModel {
	return &customStatisticsModel{
		defaultStatisticsModel: newStatisticsModel(conn, c, opts...),
	}
}

func (m *defaultStatisticsModel) FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*Statistics, error) {

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp Statistics
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
func (m *defaultStatisticsModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(statisticsRows).From(m.table)
}

func (m *defaultStatisticsModel) FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*Statistics, error) {

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*Statistics
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultStatisticsModel) FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error) {

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

func (m *defaultStatisticsModel) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("COUNT(" + field + ")").From(m.table)
}

func (m *defaultStatisticsModel) BuildFiled(old, new *Statistics) {
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
