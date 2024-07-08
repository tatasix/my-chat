package model

import (
	"chat/common/page"
	"chat/common/util"
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ ApplePaySubscriptionModel = (*customApplePaySubscriptionModel)(nil)

type (
	// ApplePaySubscriptionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customApplePaySubscriptionModel.
	ApplePaySubscriptionModel interface {
		applePaySubscriptionModel
		RowBuilder() squirrel.SelectBuilder
		CountBuilder(field string) squirrel.SelectBuilder
		FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*ApplePaySubscription, error)
		FindSum(ctx context.Context, sumBuilder squirrel.SelectBuilder) (float64, error)
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error)
		FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*ApplePaySubscription, error)
		FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, p, s int) ([]*ApplePaySubscription, error)
		BuildFiled(old, new *ApplePaySubscription)
	}

	customApplePaySubscriptionModel struct {
		*defaultApplePaySubscriptionModel
	}
)

// NewApplePaySubscriptionModel returns a model for the database table.
func NewApplePaySubscriptionModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ApplePaySubscriptionModel {
	return &customApplePaySubscriptionModel{
		defaultApplePaySubscriptionModel: newApplePaySubscriptionModel(conn, c, opts...),
	}
}

func (m *defaultApplePaySubscriptionModel) FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*ApplePaySubscription, error) {

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp ApplePaySubscription
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &resp, nil
}

func (m *defaultApplePaySubscriptionModel) FindSum(ctx context.Context, sumBuilder squirrel.SelectBuilder) (float64, error) {

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

func (m *defaultApplePaySubscriptionModel) FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error) {

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

func (m *defaultApplePaySubscriptionModel) FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*ApplePaySubscription, error) {

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*ApplePaySubscription
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultApplePaySubscriptionModel) FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, p, s int) ([]*ApplePaySubscription, error) {

	pg := page.NewPage(p, s)

	query, values, err := rowBuilder.Limit(pg.Limit()).Offset(pg.Offset()).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*ApplePaySubscription
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

// export logic
func (m *defaultApplePaySubscriptionModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(applePaySubscriptionRows).From(m.table).Where(squirrel.Eq{"is_deleted": IsNotDeleted})
}

// export logic
func (m *defaultApplePaySubscriptionModel) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("COUNT(" + field + ")").From(m.table).Where(squirrel.Eq{"is_deleted": IsNotDeleted})
}

func (m *defaultApplePaySubscriptionModel) BuildFiled(old, new *ApplePaySubscription) {
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
