package model

import (
	"chat/common/util"
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"golang.org/x/net/context"
	"time"
)

var _ WechatUserModel = (*customWechatUserModel)(nil)

type (
	// WechatUserModel is an interface to be customized, add more methods here,
	// and implement the added methods in customWechatUserModel.
	WechatUserModel interface {
		wechatUserModel
		FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*WechatUser, error)
		RowBuilder() squirrel.SelectBuilder
		FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*WechatUser, error)
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error)
		CountBuilder(field string) squirrel.SelectBuilder
		BuildFiled(old, new *WechatUser)
		UpdateLevelById(ctx context.Context, id, level int64, levelExpire string) (err error)
		UpdateLevel(ctx context.Context, expire string) (err error)
		DeleteByUser(ctx context.Context, user string) (err error)
	}

	customWechatUserModel struct {
		*defaultWechatUserModel
	}
)

// NewWechatUserModel returns a model for the database table.
func NewWechatUserModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) WechatUserModel {
	return &customWechatUserModel{
		defaultWechatUserModel: newWechatUserModel(conn, c, opts...),
	}
}

func (m *defaultWechatUserModel) FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*WechatUser, error) {

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp WechatUser
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	resp.Mobile = util.Decrypt(resp.Mobile)
	return &resp, nil
}

// export logic
func (m *defaultWechatUserModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(wechatUserRows).From(m.table).Where(squirrel.Eq{"is_deleted": IsNotDeleted})
}

func (m *defaultWechatUserModel) FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*WechatUser, error) {

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*WechatUser
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultWechatUserModel) FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error) {

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

func (m *defaultWechatUserModel) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("COUNT(" + field + ")").From(m.table)
}

func (m *defaultWechatUserModel) BuildFiled(old, new *WechatUser) {
	if new == nil {
		return
	}
	if old == nil && new.Id == 0 {
		new.Id = util.GenerateSnowflakeInt64()
		new.Mobile = util.Encrypt(new.Mobile)
		return
	}
	new.Id = old.Id
	//还需要处理
	if new.Nickname == "" {
		new.Nickname = old.Nickname
	}

	if new.UserType == 0 {
		new.UserType = old.UserType
	}

	if new.User == "" {
		new.User = old.User
	}

	if new.Avatar == "" {
		new.Avatar = old.Avatar
	}

	if new.Unionid == "" {
		new.Unionid = old.Unionid
	}

	if new.Openid == "" {
		new.Openid = old.Openid
	}

	if new.Province == "" {
		new.Province = old.Province
	}

	if new.Country == "" {
		new.Country = old.Country
	}

	if new.City == "" {
		new.City = old.City
	}

	if new.Gender == 0 {
		new.Gender = old.Gender
	}

	if new.Birthday == "" {
		new.Birthday = old.Birthday
	}

	if new.Constellation == "" {
		new.Constellation = old.Constellation
	}

	if new.Salt == "" {
		new.Salt = old.Salt
	}

	if new.Name == "" {
		new.Name = old.Name
	}

	if new.Mobile == "" {
		new.Mobile = util.Encrypt(old.Mobile)
	} else {
		new.Mobile = util.Encrypt(new.Mobile)
	}

	if new.Status == 0 {
		new.Status = old.Status
	}

	if new.Password == "" {
		new.Password = old.Password
	}

	if new.Level == 0 {
		new.Level = old.Level
	}

	if new.LevelExpire == "" {
		new.LevelExpire = old.LevelExpire
	}

	if new.IsLoginFree == 0 {
		new.IsLoginFree = old.IsLoginFree
	}
	if new.Mbti == "" {
		new.Mbti = old.Mbti
	}
	new.UpdatedAt = time.Now()

	return
}

func (m *defaultWechatUserModel) UpdateLevelById(ctx context.Context, id, level int64, levelExpire string) (err error) {

	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set `level`= ?,`level_expire`= ? where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, level, levelExpire, id)
	})

	return
}

func (m *defaultWechatUserModel) UpdateLevel(ctx context.Context, expire string) (err error) {

	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set `level`= 0 where `level_expire` <= ?", m.table)
		return conn.ExecCtx(ctx, query, expire)
	})

	return
}

func (m *defaultWechatUserModel) DeleteByUser(ctx context.Context, user string) (err error) {

	_, err = m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `user` = ?", m.table)
		return conn.ExecCtx(ctx, query, user)
	})

	return
}
