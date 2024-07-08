package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ QuestionnaireQuestionModel = (*customQuestionnaireQuestionModel)(nil)

type (
	// QuestionnaireQuestionModel is an interface to be customized, add more methods here,
	// and implement the added methods in customQuestionnaireQuestionModel.
	QuestionnaireQuestionModel interface {
		questionnaireQuestionModel
		FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*QuestionnaireQuestion, error)
		RowBuilder() squirrel.SelectBuilder
		FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*QuestionnaireQuestion, error)
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error)
		CountBuilder(field string) squirrel.SelectBuilder
		UpdateTimesAndToken(ctx context.Context, id, times, token int64) (err error)
	}

	customQuestionnaireQuestionModel struct {
		*defaultQuestionnaireQuestionModel
	}
)

// NewQuestionnaireQuestionModel returns a model for the database table.
func NewQuestionnaireQuestionModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) QuestionnaireQuestionModel {
	return &customQuestionnaireQuestionModel{
		defaultQuestionnaireQuestionModel: newQuestionnaireQuestionModel(conn, c, opts...),
	}
}

func (m *defaultQuestionnaireQuestionModel) FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*QuestionnaireQuestion, error) {

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp QuestionnaireQuestion
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
func (m *defaultQuestionnaireQuestionModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(questionnaireQuestionRows).From(m.table)
}

func (m *defaultQuestionnaireQuestionModel) FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*QuestionnaireQuestion, error) {

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*QuestionnaireQuestion
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultQuestionnaireQuestionModel) FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error) {

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

func (m *defaultQuestionnaireQuestionModel) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("COUNT(" + field + ")").From(m.table)
}

func (m *defaultQuestionnaireQuestionModel) UpdateTimesAndToken(ctx context.Context, id, times, token int64) (err error) {

	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set `times`= `times` + ? ,`token`= `token` + ? where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, times, token, id)
	})

	return
}

type QuestionnaireQuestionOption struct {
	Id      int64  `json:"id"`
	Content string `json:"content"`
	Score   int64  `json:"score"`
	Mbti    string `json:"mbti"`
}
