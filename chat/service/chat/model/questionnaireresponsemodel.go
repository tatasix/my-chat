package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ QuestionnaireResponseModel = (*customQuestionnaireResponseModel)(nil)

type (
	// QuestionnaireResponseModel is an interface to be customized, add more methods here,
	// and implement the added methods in customQuestionnaireResponseModel.
	QuestionnaireResponseModel interface {
		questionnaireResponseModel
		FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*QuestionnaireResponse, error)
		RowBuilder() squirrel.SelectBuilder
		FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*QuestionnaireResponse, error)
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error)
		CountBuilder(field string) squirrel.SelectBuilder
		UpdateTimesAndToken(ctx context.Context, id, times, token int64) (err error)
		FindSum(ctx context.Context, sumBuilder squirrel.SelectBuilder) (int64, error)
		SumBuilder(field string) squirrel.SelectBuilder
	}

	customQuestionnaireResponseModel struct {
		*defaultQuestionnaireResponseModel
	}
)

// NewQuestionnaireResponseModel returns a model for the database table.
func NewQuestionnaireResponseModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) QuestionnaireResponseModel {
	return &customQuestionnaireResponseModel{
		defaultQuestionnaireResponseModel: newQuestionnaireResponseModel(conn, c, opts...),
	}
}

func (m *defaultQuestionnaireResponseModel) FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*QuestionnaireResponse, error) {

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp QuestionnaireResponse
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
func (m *defaultQuestionnaireResponseModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(questionnaireResponseRows).From(m.table)
}

func (m *defaultQuestionnaireResponseModel) FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*QuestionnaireResponse, error) {

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*QuestionnaireResponse
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultQuestionnaireResponseModel) FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error) {

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

func (m *defaultQuestionnaireResponseModel) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("COUNT(" + field + ")").From(m.table)
}

func (m *defaultQuestionnaireResponseModel) UpdateTimesAndToken(ctx context.Context, id, times, token int64) (err error) {

	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set `times`= `times` + ? ,`token`= `token` + ? where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, times, token, id)
	})

	return
}

func (m *defaultQuestionnaireResponseModel) FindSum(ctx context.Context, sumBuilder squirrel.SelectBuilder) (int64, error) {

	query, values, err := sumBuilder.ToSql()
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

func (m *defaultQuestionnaireResponseModel) SumBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("IFNULL(SUM(" + field + "),0)").From(m.table)
}
