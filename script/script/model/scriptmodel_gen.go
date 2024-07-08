// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	scriptFieldNames          = builder.RawFieldNames(&Script{})
	scriptRows                = strings.Join(scriptFieldNames, ",")
	scriptRowsExpectAutoSet   = strings.Join(stringx.Remove(scriptFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	scriptRowsWithPlaceHolder = strings.Join(stringx.Remove(scriptFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheScriptIdPrefix = "cache:script:id:"
)

type (
	scriptModel interface {
		Insert(ctx context.Context, data *Script) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Script, error)
		Update(ctx context.Context, data *Script) error
		Delete(ctx context.Context, id int64) error
	}

	defaultScriptModel struct {
		sqlc.CachedConn
		table string
	}

	Script struct {
		Id         int64     `db:"id"`
		Name       string    `db:"name"`        // 脚本名称
		Path       string    `db:"path"`        // 脚本保存路径
		ScriptType string    `db:"script_type"` // 脚本类型
		IsDelete   int64     `db:"is_delete"`   // 是否删除，0否，1是
		IsEnable   int64     `db:"is_enable"`   // 是否启用，0否，1是
		CreatedAt  time.Time `db:"created_at"`  // 创建时间
		UpdatedAt  time.Time `db:"updated_at"`  // 更新时间
	}
)

func newScriptModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultScriptModel {
	return &defaultScriptModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`script`",
	}
}

func (m *defaultScriptModel) Delete(ctx context.Context, id int64) error {
	scriptIdKey := fmt.Sprintf("%s%v", cacheScriptIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, scriptIdKey)
	return err
}

func (m *defaultScriptModel) FindOne(ctx context.Context, id int64) (*Script, error) {
	scriptIdKey := fmt.Sprintf("%s%v", cacheScriptIdPrefix, id)
	var resp Script
	err := m.QueryRowCtx(ctx, &resp, scriptIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", scriptRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultScriptModel) Insert(ctx context.Context, data *Script) (sql.Result, error) {
	scriptIdKey := fmt.Sprintf("%s%v", cacheScriptIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?, ?)", m.table, scriptRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Name, data.Path, data.ScriptType, data.IsDelete, data.IsEnable)
	}, scriptIdKey)
	return ret, err
}

func (m *defaultScriptModel) Update(ctx context.Context, data *Script) error {
	scriptIdKey := fmt.Sprintf("%s%v", cacheScriptIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, scriptRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.Name, data.Path, data.ScriptType, data.IsDelete, data.IsEnable, data.Id)
	}, scriptIdKey)
	return err
}

func (m *defaultScriptModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheScriptIdPrefix, primary)
}

func (m *defaultScriptModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", scriptRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultScriptModel) tableName() string {
	return m.table
}