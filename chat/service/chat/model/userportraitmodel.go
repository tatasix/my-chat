package model

import (
	"chat/common/util"
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserPortraitModel = (*customUserPortraitModel)(nil)

type (
	// UserPortraitModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserPortraitModel.
	UserPortraitModel interface {
		userPortraitModel
		FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*UserPortrait, error)
		RowBuilder() squirrel.SelectBuilder
		FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*UserPortrait, error)
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error)
		CountBuilder(field string) squirrel.SelectBuilder
		BuildFiled(old, new *UserPortrait)
	}

	customUserPortraitModel struct {
		*defaultUserPortraitModel
	}
)

// NewUserPortraitModel returns a model for the database table.
func NewUserPortraitModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) UserPortraitModel {
	return &customUserPortraitModel{
		defaultUserPortraitModel: newUserPortraitModel(conn, c, opts...),
	}
}

func (m *defaultUserPortraitModel) FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*UserPortrait, error) {

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp UserPortrait
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
func (m *defaultUserPortraitModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(userPortraitRows).From(m.table)
}

func (m *defaultUserPortraitModel) FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*UserPortrait, error) {

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*UserPortrait
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultUserPortraitModel) FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error) {

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

func (m *defaultUserPortraitModel) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("COUNT(" + field + ")").From(m.table)
}

func (m *defaultUserPortraitModel) BuildFiled(old, new *UserPortrait) {
	if new == nil {
		return
	}
	if old == nil && new.Id == 0 {
		new.Id = util.GenerateSnowflakeInt64()
		return
	}
	new.Id = old.Id

	return
}
