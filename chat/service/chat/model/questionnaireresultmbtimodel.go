package model

import (
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ QuestionnaireResultMbtiModel = (*customQuestionnaireResultMbtiModel)(nil)

type (
	// QuestionnaireResultMbtiModel is an interface to be customized, add more methods here,
	// and implement the added methods in customQuestionnaireResultMbtiModel.
	QuestionnaireResultMbtiModel interface {
		questionnaireResultMbtiModel
		FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*QuestionnaireResultMbti, error)
		RowBuilder() squirrel.SelectBuilder
		FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*QuestionnaireResultMbti, error)
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error)
	}

	customQuestionnaireResultMbtiModel struct {
		*defaultQuestionnaireResultMbtiModel
	}
)

// NewQuestionnaireResultMbtiModel returns a model for the database table.
func NewQuestionnaireResultMbtiModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) QuestionnaireResultMbtiModel {
	return &customQuestionnaireResultMbtiModel{
		defaultQuestionnaireResultMbtiModel: newQuestionnaireResultMbtiModel(conn, c, opts...),
	}
}

func (m *defaultQuestionnaireResultMbtiModel) FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*QuestionnaireResultMbti, error) {

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp QuestionnaireResultMbti
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
func (m *defaultQuestionnaireResultMbtiModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(questionnaireResultMbtiRows).From(m.table)
}

func (m *defaultQuestionnaireResultMbtiModel) FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*QuestionnaireResultMbti, error) {

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*QuestionnaireResultMbti
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultQuestionnaireResultMbtiModel) FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error) {

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

type QuestionnaireResultMbtiScore struct {
	Type        int64   `json:"type"`
	Score       float64 `json:"score"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
}
