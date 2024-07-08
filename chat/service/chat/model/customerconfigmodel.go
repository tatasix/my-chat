package model

import (
	"chat/common/util"
	"context"
	"github.com/Masterminds/squirrel"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ CustomerConfigModel = (*customCustomerConfigModel)(nil)

type (
	// CustomerConfigModel is an interface to be customized, add more methods here,
	// and implement the added methods in customCustomerConfigModel.
	CustomerConfigModel interface {
		customerConfigModel
		FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*CustomerConfig, error)
		RowBuilder() squirrel.SelectBuilder
		FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*CustomerConfig, error)
		CountBuilder(field string) squirrel.SelectBuilder
		FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error)
		BuildFiled(old, new *CustomerConfig)
	}

	customCustomerConfigModel struct {
		*defaultCustomerConfigModel
	}

	PromptStates struct {
		Id       int32   `json:"id"`
		Name     string  `json:"name"`
		NextNode []int32 `json:"next_node"`
	}

	PayConfig struct {
		Enable   bool   `json:"enable"`
		Amount   int64  `json:"amount"`
		Type     int64  `json:"type"`
		Times    int64  `json:"times"`
		Describe string `json:"describe"`
		Period   int64  `json:"period"` //计数周期，1 按天；2 没有时间限制
	}
)

// NewCustomerConfigModel returns a model for the database table.
func NewCustomerConfigModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) CustomerConfigModel {
	return &customCustomerConfigModel{
		defaultCustomerConfigModel: newCustomerConfigModel(conn, c, opts...),
	}
}

func (m *defaultCustomerConfigModel) FindOneByQuery(ctx context.Context, rowBuilder squirrel.SelectBuilder) (*CustomerConfig, error) {

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp CustomerConfig
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
func (m *defaultCustomerConfigModel) RowBuilder() squirrel.SelectBuilder {
	return squirrel.Select(customerConfigRows).From(m.table).Where(squirrel.Eq{"is_deleted": IsNotDeleted})
}

func (m *defaultCustomerConfigModel) FindAll(ctx context.Context, rowBuilder squirrel.SelectBuilder) ([]*CustomerConfig, error) {

	query, values, err := rowBuilder.ToSql()
	if err != nil {
		return nil, err
	}

	var resp []*CustomerConfig
	err = m.QueryRowsNoCacheCtx(ctx, &resp, query, values...)
	switch err {
	case nil:
		return resp, nil
	default:
		return nil, err
	}
}

func (m *defaultCustomerConfigModel) CountBuilder(field string) squirrel.SelectBuilder {
	return squirrel.Select("COUNT(" + field + ")").From(m.table).Where(squirrel.Eq{"is_deleted": IsNotDeleted})
}

func (m *defaultCustomerConfigModel) FindCount(ctx context.Context, countBuilder squirrel.SelectBuilder) (int64, error) {

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

func (m *defaultCustomerConfigModel) BuildFiled(old, new *CustomerConfig) {

	if new == nil {
		return
	}
	if old == nil && new.Id == 0 {
		new.Id = util.GenerateSnowflakeInt64()
		return
	}
	new.Id = old.Id

	if new.Pay == "" {
		new.Pay = old.Pay
	}

	if new.PromptStates == "" {
		new.PromptStates = old.PromptStates
	}

	if new.Multiple == 0 {
		new.Multiple = old.Multiple
	}

	return
}

type CustomerConfigConfig struct {
	Method       int64                            `json:"method"`
	Total        int64                            `json:"total"`
	Mod          CustomerConfigConfigMod          `json:"mod"`
	Result       []CustomerConfigConfigResultType `json:"result"`
	Mbti93Result []Mbti93                         `json:"mbti93_result"`
}
type CustomerConfigConfigMod struct {
	Type int64   `json:"type"`
	Rate float32 `json:"rate"`
}
type CustomerConfigConfigResultType struct {
	Type        int64                        `json:"type"`
	Total       float64                      `json:"total"`
	Name        string                       `json:"name"`
	Description string                       `json:"description"`
	Result      []CustomerConfigConfigResult `json:"result"`
}

type CustomerConfigConfigResult struct {
	BaseResult
	MbtiResult
}

type BaseResult struct {
	Min float64 `json:"min"`
	Max float64 `json:"max"`
	Res string  `json:"res"` // 结论
}

type MbtiResult struct {
	Name         string `json:"name"`      // 名称
	Nickname     string `json:"nickname"`  // 昵称
	SimpleResult string `json:"res1"`      // 一句话结果
	NameType     string `json:"name_type"` // 类型
	Label        string `json:"label"`     // 标签
	Partner      string `json:"partner"`   // 拍档
	Mate         string `json:"mate"`      // 朋友
	Content      string `json:"content"`   // ai 生成内容
	PcImage      string `json:"pc_image"`  // pc 图片
	H5Image      string `json:"h5_image"`  // h5 图片
	Color        int64  `json:"color"`     // 颜色

}

type Mbti93 struct {
	NameType string `json:"name_type"`
	Name     string `json:"name"`
	Mate     string `json:"mate"`
	Partner  string `json:"partner"`
	Color    int64  `json:"color"`
	Label    string `json:"label"`
	PcImage  string `json:"pc_image"`
	H5Image  string `json:"h5_image"`
	Nickname string `json:"nickname"`
	Res1     string `json:"res1"`
	Res      string `json:"res"`
}
