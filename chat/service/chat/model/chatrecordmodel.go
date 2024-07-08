package model

import (
	"bytes"
	"chat/common/page"
	"chat/common/util"
	"context"
	"database/sql"
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"log"
)

var _ ChatRecordModel = (*customChatRecordModel)(nil)

type (
	// ChatRecordModel is an interface to be customized, add more methods here,
	// and implement the added methods in customChatRecordModel.
	ChatRecordModel interface {
		chatRecordModel
		Trans(ctx context.Context, fn func(context context.Context, session sqlx.Session) error) error
		RowBuilder() squirrel.SelectBuilder
		CountBuilder(field string) squirrel.SelectBuilder
		SumBuilder(field string) squirrel.SelectBuilder
		FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*ChatRecord, error)
		FindSum(ctx context.Context, sumBuilder squirrel.SelectBuilder) (float64, error)
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error)
		FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*ChatRecord, error)
		FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, p, s int) ([]*ChatRecord, error)
		BuildFiled(old, new *ChatRecord)
		BatchInsert(ctx context.Context, data []*ChatRecord) (sql.Result, error)
	}

	customChatRecordModel struct {
		*defaultChatRecordModel
	}

	ExportChatRecord struct {
		Id               int64
		RelationId       int64  // 关联id
		UserName         string // weCom用户标识/customer用户标识
		MessageId        string // message_id customer消息唯一ID
		OpenKfName       string // 客服标识
		AgentName        string // 应用ID
		Content          string // 消息内容
		Emoji            int64  // emoji类型
		ChatType         int64  // 聊天类型
		AnswerOrQuestion int64  // 问题还是答案:1问题，2答案
		MessageType      int64  // 消息类型
		CreatedAt        string // 创建时间
		UpdatedAt        string // 更新时间
	}
)

// NewChatRecordModel returns a model for the database table.
func NewChatRecordModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) ChatRecordModel {
	return &customChatRecordModel{
		defaultChatRecordModel: newChatRecordModel(conn, c, opts...),
	}
}

func (m *defaultChatRecordModel) FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*ChatRecord, error) {

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp ChatRecord
	err = m.QueryRowNoCacheCtx(ctx, &resp, query, values...)
	if err != nil {
		if err == sqlx.ErrNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &resp, nil
}

func (m *defaultChatRecordModel) FindSum(ctx context.Context, sumBuilder squirrel.SelectBuilder) (float64, error) {

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

func (m *defaultChatRecordModel) FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error) {

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

func (m *defaultChatRecordModel) FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*ChatRecord, error) {

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*ChatRecord
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultChatRecordModel) FindPageListByPage(ctx context.Context, rowBuilder squirrel.SelectBuilder, p, s int) ([]*ChatRecord, error) {

	pg := page.NewPage(p, s)

	query, values, err := rowBuilder.Limit(pg.Limit()).Offset(pg.Offset()).ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*ChatRecord
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

// export logic
func (m *defaultChatRecordModel) Trans(ctx context.Context, fn func(ctx context.Context, session sqlx.Session) error) error {

	return m.TransactCtx(ctx, func(ctx context.Context, session sqlx.Session) error {
		return fn(ctx, session)
	})

}

// export logic
func (m *defaultChatRecordModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(chatRecordRows).From(m.table)
}

// export logic
func (m *defaultChatRecordModel) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("COUNT(" + field + ")").From(m.table)
}

// export logic
func (m *defaultChatRecordModel) SumBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("IFNULL(SUM(" + field + "),0)").From(m.table)
}

func (m *defaultChatRecordModel) BuildFiled(old, new *ChatRecord) {
	if new == nil {
		return
	}
	if old == nil {
		newId := util.GenerateSnowflakeInt64()
		if new.Id == 0 {
			new.Id = newId
		}
		if new.RelationId == 0 {
			new.RelationId = newId
		}
		return
	}
	new.Id = old.Id
	return
}

func (m *defaultChatRecordModel) BatchInsert(ctx context.Context, data []*ChatRecord) (sql.Result, error) {

	var placeholders bytes.Buffer
	var values []interface{}

	for i, record := range data {
		if i > 0 {
			placeholders.WriteString(", ")
		}
		placeholders.WriteString("(?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")

		values = append(values, record.Id, record.RelationId, record.User, record.MessageId, record.OpenKfId, record.AgentId, record.Content, record.Emoji, record.ChatType, record.AnswerOrQuestion, record.MessageType, util.TimeFormat(record.CreatedAt), util.TimeFormat(record.UpdatedAt))
	}

	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("INSERT INTO %s (%s) VALUES %s ", m.table, chatRecordRowsMigragteExpectAutoSet, placeholders.String())
		stmt, err := conn.Prepare(query)
		if err != nil {
			log.Fatal(err)
		}
		defer stmt.Close()

		return stmt.Exec(values...)

	}, "")
	return ret, err
}
