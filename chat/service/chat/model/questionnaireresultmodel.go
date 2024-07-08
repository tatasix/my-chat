package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ QuestionnaireResultModel = (*customQuestionnaireResultModel)(nil)

type (
	// QuestionnaireResultModel is an interface to be customized, add more methods here,
	// and implement the added methods in customQuestionnaireResultModel.
	QuestionnaireResultModel interface {
		questionnaireResultModel
		FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*QuestionnaireResult, error)
		RowBuilder() squirrel.SelectBuilder
		FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*QuestionnaireResult, error)
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error)
		CountBuilder(field string) squirrel.SelectBuilder
		UpdateTimesAndToken(ctx context.Context, id, times, token int64) (err error)
	}

	customQuestionnaireResultModel struct {
		*defaultQuestionnaireResultModel
	}
)

// NewQuestionnaireResultModel returns a model for the database table.
func NewQuestionnaireResultModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) QuestionnaireResultModel {
	return &customQuestionnaireResultModel{
		defaultQuestionnaireResultModel: newQuestionnaireResultModel(conn, c, opts...),
	}
}

func (m *defaultQuestionnaireResultModel) FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*QuestionnaireResult, error) {

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp QuestionnaireResult
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
func (m *defaultQuestionnaireResultModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(questionnaireResultRows).From(m.table)
}

func (m *defaultQuestionnaireResultModel) FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*QuestionnaireResult, error) {

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*QuestionnaireResult
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultQuestionnaireResultModel) FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error) {

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

func (m *defaultQuestionnaireResultModel) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("COUNT(" + field + ")").From(m.table)
}

func (m *defaultQuestionnaireResultModel) UpdateTimesAndToken(ctx context.Context, id, times, token int64) (err error) {

	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set `times`= `times` + ? ,`token`= `token` + ? where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, times, token, id)
	})

	return
}
