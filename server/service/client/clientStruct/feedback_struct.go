package clientStruct

type GetFeedbackListRequest struct {
	Page           int    `json:"page"`      // 页码
	PageSize       int    `json:"page_size"` // 每页大小
	StartCreatedAt string `json:"start_created_at"`
	EndCreatedAt   string `json:"end_created_at"`
	User           string `json:"user"`
	OpenKfId       string `json:"open_kf_id"`
	Status         int64  `json:"status"` // 状态:1待处理，2已处理
}

type GetFeedbackResponse struct {
	List     []FeedbackResponse `json:"list"`
	Total    int64              `json:"total"`
	Page     int                `json:"page"`
	PageSize int                `json:"pageSize"`
}

type FeedbackResponse struct {
	Id        int64  `json:"id"`         // weCom用户标识/customer用户标识
	User      string `json:"user"`       // weCom用户标识/customer用户标识
	MessageId string `json:"message_id"` // message_id customer消息唯一ID
	OpenKfId  string `json:"open_kf_id"` // 客服标识
	Title     string `json:"title"`      // 用户反馈标题
	Content   string `json:"content"`    // 用户反馈内容
	Reply     string `json:"reply"`      // 回复内容
	Status    int64  `json:"status"`     // 状态:1待处理，2已处理
	CreatedAt string `json:"created_at"` // 创建时间
	UpdatedAt string `json:"updated_at"` // 更新时间
	IsDeleted int64  `json:"is_deleted"` // 是否删除，0否，1是
}

type FindFeedbackRequest struct {
	Id int64 `json:"id"`
}
