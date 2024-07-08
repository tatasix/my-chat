// 自动生成模板Config
package chatAdmin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Config 结构体
type Config struct {
	global.GVA_MODEL
	ConfigType  *int   `json:"configType" form:"configType" gorm:"column:config_type;comment:配置类型：1，openai key;"`
	Name        string `json:"name" form:"name" gorm:"column:name;comment:配置名;size:64;"`
	Description string `json:"description" form:"description" gorm:"column:description;comment:配置描述;size:256;"`
	Value       string `json:"value" form:"value" gorm:"column:value;comment:配置内容;"`
	CreatedBy   string `json:"createdBy" form:"createdBy" gorm:"column:created_by;comment:创建人;size:64;"`
	UpdatedBy   string `json:"updatedBy" form:"updatedBy" gorm:"column:updated_by;comment:更新人;size:64;"`
}

// TableName Config 表名
func (Config) TableName() string {
	return "config"
}
