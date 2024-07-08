// 自动生成模板Feedback
package chatAdmin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
)

// Feedback 结构体
type Feedback struct {
	global.GVA_MODEL
	User      string `json:"user" form:"user" gorm:"column:user;comment:weCom用户标识/customer用户标识;size:128;"`
	MessageId string `json:"messageId" form:"messageId" gorm:"column:message_id;comment:message_id customer消息唯一ID;size:128;"`
	OpenKfId  string `json:"openKfId" form:"openKfId" gorm:"column:open_kf_id;comment:客服标识;size:128;"`
	Title     string `json:"title" form:"title" gorm:"column:title;comment:用户反馈标题;size:500;"`
	Content   string `json:"content" form:"content" gorm:"column:content;comment:用户反馈内容;"`
	Reply     string `json:"reply" form:"reply" gorm:"column:reply;comment:回复内容;"`
	Status    int64  `json:"status" form:"status" gorm:"column:status;comment:状态:1待处理，2已处理;"`
	IsDeleted *bool  `json:"isDeleted" form:"isDeleted" gorm:"column:is_deleted;comment:是否删除，0否，1是;"`
}

// TableName Feedback 表名
func (Feedback) TableName() string {
	return "feedback"
}
