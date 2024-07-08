package client

import (
	"encoding/json"
	"github.com/flipped-aurora/gin-vue-admin/server/model/chatAdmin"
	chatAdminReq "github.com/flipped-aurora/gin-vue-admin/server/model/chatAdmin/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/service/client/assembler"
	"github.com/flipped-aurora/gin-vue-admin/server/service/client/clientStruct"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"github.com/flipped-aurora/gin-vue-admin/server/vars"
)

type FeedbackService struct {
}

func (c *FeedbackService) CreateFeedback(Feedback *chatAdmin.Feedback) (err error) {

	return
}

func (c *FeedbackService) DeleteFeedback(Feedback chatAdmin.Feedback) (err error) {

	return
}

func (c *FeedbackService) DeleteFeedbackByIds(ids request.IdsReq) (err error) {

	return
}

func (c *FeedbackService) UpdateFeedback(Feedback chatAdmin.Feedback) (err error) {

	return
}

func (c *FeedbackService) GetFeedback(id uint) (Feedback chatAdmin.Feedback, err error) {
	param := clientStruct.FindFeedbackRequest{
		Id: int64(id),
	}

	jsonParam, _ := json.Marshal(param)

	result, err := utils.Post(utils.GetChatServerHost()+vars.ChatFindFeedbackUri, jsonParam, nil)

	if err != nil {
		return
	}
	var resultInfo clientStruct.FeedbackResponse
	err = json.Unmarshal(result, &resultInfo)
	if err != nil {
		return
	}
	if resultInfo.Id <= 0 {
		return
	}
	Feedback = assembler.POTODTOGetFeedback(resultInfo)
	return
}

func (c *FeedbackService) GetFeedbackInfoList(info chatAdminReq.FeedbackSearch) (list []chatAdmin.Feedback, total int64, err error) {

	param := clientStruct.GetFeedbackListRequest{
		Page:           info.Page,
		PageSize:       info.PageSize,
		User:           info.User,
		OpenKfId:       info.OpenKfId,
		Status:         info.Status,
		StartCreatedAt: utils.TimeToString(info.StartCreatedAt),
		EndCreatedAt:   utils.TimeToString(info.EndCreatedAt),
	}

	jsonParam, _ := json.Marshal(param)

	result, err := utils.Post(utils.GetChatServerHost()+vars.ChatGetFeedbackUri, jsonParam, nil)

	if err != nil {
		return
	}
	var resultInfo clientStruct.GetFeedbackResponse
	err = json.Unmarshal(result, &resultInfo)
	if err != nil {
		return
	}
	if resultInfo.List == nil || len(resultInfo.List) <= 0 {
		return
	}
	list = assembler.POTODTOGetFeedbackList(resultInfo.List)
	total = resultInfo.Total
	return
}
