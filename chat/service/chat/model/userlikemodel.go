package model

import (
	"chat/common/page"
	"chat/common/util"
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ UserLikeModel = (*customUserLikeModel)(nil)

type (
	// UserLikeModel is an interface to be customized, add more methods here,
	// and implement the added methods in customUserLikeModel.
	UserLikeModel interface {
		userLikeModel
		Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error
		RowBuilder() squirrel.SelectBuilder
		CountBuilder(field string) squirrel.SelectBuilder
		SumBuilder(field string) squirrel.SelectBuilder
		FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*UserLike, error)
		FindSum(ctx context.Context, sumBuilder squirrel.SelectBuilder) (float64, error)
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error)
		FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*UserLike, error)
		FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, p, s int) ([]*UserLike, error)
		BuildFiled(old, new *UserLike)
	}

	customUserLikeModel struct {
		*defaultUserLikeModel
	}
)

// NewUserLikeModel returns a model for the database table.
func NewUserLikeModel(conn sqlx.SqlConn, c cache.CacheConf) UserLikeModel {
	return &customUserLikeModel{
		defaultUserLikeModel: newUserLikeModel(conn, c),
	}
}

func (m *defaultUserLikeModel) FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*UserLike, error) {

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp UserLike
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

func (m *defaultUserLikeModel) FindSum(ctx context.Context, sumBuilder squirrel.SelectBuilder) (float64, error) {

	query, values, err := sumBuilder.ToSql()
	if err != nil {
		return 0, err
	}

	var resp float64
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}

func (m *defaultUserLikeModel) FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error) {

	query, values, err := countBuilder.ToSql()
	if err != nil {
		return 0, err
	}

	var resp int64
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return 0, err
	}
}

func (m *defaultUserLikeModel) FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*UserLike, error) {

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*UserLike
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultUserLikeModel) FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, p, s int) ([]*UserLike, error) {

	pg := page.NewPage(p, s)

	query, values, err := rowBuilder.Limit(pg.Limit()).Offset(pg.Offset()).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*UserLike
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

// export logic
func (m *defaultUserLikeModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {

	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})

}

// export logic
func (m *defaultUserLikeModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(userLikeRows).From(m.table)
}

// export logic
func (m *defaultUserLikeModel) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("COUNT(" + field + ")").From(m.table)
}

// export logic
func (m *defaultUserLikeModel) SumBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("IFNULL(SUM(" + field + "),0)").From(m.table)
}

func (m *defaultUserLikeModel) BuildFiled(old, new *UserLike) {
	if new == nil {
		return
	}
	if old == nil {
		newId := util.GenerateSnowflakeInt64()
		if new.Id == 0 {
			new.Id = newId
		}
		return
	}
	new.Id = old.Id
	return
}
