package clientStruct

type IdsRequest struct {
	Ids []int64 `json:"id"`
}
type Response struct {
	Message string `json:"message"`
}

type IdRequest struct {
	Id int64 `json:"id"`
}

type PageInfo struct {
	Page     int    `json:"page" form:"page"`         // 页码
	PageSize int    `json:"pageSize" form:"pageSize"` // 每页大小
	Keyword  string `json:"keyword" form:"keyword"`   //关键字
}

type IdV2Request struct {
	Id       int64  `json:"id"`
	UpdateBy string `json:"updated_by"` // 修改人
}
