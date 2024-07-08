package model

import (
	"chat/common/page"
	"chat/common/util"
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ FeedbackModel = (*customFeedbackModel)(nil)

type (
	// FeedbackModel is an interface to be customized, add more methods here,
	// and implement the added methods in customFeedbackModel.
	FeedbackModel interface {
		feedbackModel
		RowBuilder() squirrel.SelectBuilder
		CountBuilder(field string) squirrel.SelectBuilder
		FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*Feedback, error)
		FindSum(ctx context.Context, sumBuilder squirrel.SelectBuilder) (float64, error)
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error)
		FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*Feedback, error)
		FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, p, s int) ([]*Feedback, error)
		BuildFiled(old, new *Feedback)
	}

	customFeedbackModel struct {
		*defaultFeedbackModel
	}
)

// NewFeedbackModel returns a model for the database table.
func NewFeedbackModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) FeedbackModel {
	return &customFeedbackModel{
		defaultFeedbackModel: newFeedbackModel(conn, c, opts...),
	}
}

func (m *defaultFeedbackModel) FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*Feedback, error) {

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp Feedback
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &resp, nil
}

func (m *defaultFeedbackModel) FindSum(ctx context.Context, sumBuilder squirrel.SelectBuilder) (float64, error) {

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

func (m *defaultFeedbackModel) FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error) {

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

func (m *defaultFeedbackModel) FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*Feedback, error) {

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*Feedback
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultFeedbackModel) FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, p, s int) ([]*Feedback, error) {

	pg := page.NewPage(p, s)

	query, values, err := rowBuilder.Limit(pg.Limit()).Offset(pg.Offset()).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*Feedback
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

// export logic
func (m *defaultFeedbackModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(feedbackRows).From(m.table).Where(squirrel.Eq{"is_deleted": IsNotDeleted})
}

// export logic
func (m *defaultFeedbackModel) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("COUNT(" + field + ")").From(m.table).Where(squirrel.Eq{"is_deleted": IsNotDeleted})
}

func (m *defaultFeedbackModel) BuildFiled(old, new *Feedback) {
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
