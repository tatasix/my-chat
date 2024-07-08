package model

import (
	"chat/common/util"
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"time"
)

var _ PromptModel = (*customPromptModel)(nil)

type (
	// PromptModel is an interface to be customized, add more methods here,
	// and implement the added methods in customPromptModel.
	PromptModel interface {
		promptModel
		FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*Prompt, error)
		RowBuilder() squirrel.SelectBuilder
		FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*Prompt, error)
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error)
		CountBuilder(field string) squirrel.SelectBuilder
		BuildFiled(old, new *Prompt)
	}

	customPromptModel struct {
		*defaultPromptModel
	}
)

// NewPromptModel returns a model for the database table.
func NewPromptModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) PromptModel {
	return &customPromptModel{
		defaultPromptModel: newPromptModel(conn, c, opts...),
	}
}

func (m *defaultPromptModel) FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*Prompt, error) {

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp Prompt
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
func (m *defaultPromptModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(promptRows).From(m.table).Where(squirrel.Eq{"is_deleted": IsNotDeleted})
}

func (m *defaultPromptModel) FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*Prompt, error) {

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*Prompt
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultPromptModel) FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error) {

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

func (m *defaultPromptModel) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("COUNT(" + field + ")").From(m.table)
}

func (m *defaultPromptModel) BuildFiled(old, new *Prompt) {
	if new == nil {
		return
	}
	if old == nil && new.Id == 0 {
		new.Id = util.GenerateSnowflakeInt64()
		return
	}
	new.Id = old.Id
	//还需要处理
	//if new.Nickname == "" {
	//	new.Nickname = old.Nickname
	//}
	//
	//if new.UserType == 0 {
	//	new.UserType = old.UserType
	//}
	//
	//if new.User == "" {
	//	new.User = old.User
	//}
	//
	//if new.Avatar == "" {
	//	new.Avatar = old.Avatar
	//}
	//
	//if new.Unionid == "" {
	//	new.Unionid = old.Unionid
	//}
	//
	//if new.Openid == "" {
	//	new.Openid = old.Openid
	//}
	//
	//if new.Province == "" {
	//	new.Province = old.Province
	//}
	//
	//if new.Country == "" {
	//	new.Country = old.Country
	//}
	//
	//if new.City == "" {
	//	new.City = old.City
	//}
	//
	//if new.Gender == 0 {
	//	new.Gender = old.Gender
	//}
	//
	//if new.Birthday == "" {
	//	new.Birthday = old.Birthday
	//}
	//
	//if new.Constellation == "" {
	//	new.Constellation = old.Constellation
	//}
	//
	//if new.Salt == "" {
	//	new.Salt = old.Salt
	//}
	//
	//if new.Name == "" {
	//	new.Name = old.Name
	//}
	//
	//if new.Mobile == "" {
	//	new.Mobile = old.Mobile
	//}
	//
	//if new.Status == 0 {
	//	new.Status = old.Status
	//}
	//
	//if new.Password == "" {
	//	new.Password = old.Password
	//}
	//
	//if new.IsLoginFree == 0 {
	//	new.IsLoginFree = old.IsLoginFree
	//}

	new.UpdatedAt = time.Now()

	return
}
