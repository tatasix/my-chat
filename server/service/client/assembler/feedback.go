package assembler

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/chatAdmin"
	"github.com/flipped-aurora/gin-vue-admin/server/service/client/clientStruct"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
)

func POTODTOGetFeedbackList(feedbackPos []clientStruct.FeedbackResponse) (dto []chatAdmin.Feedback) {
	if len(feedbackPos) <= 0 {
		return
	}

	for _, v := range feedbackPos {
		dto = append(dto, POTODTOGetFeedback(v))
	}
	return
}

func POTODTOGetFeedback(feedbackPo clientStruct.FeedbackResponse) (dto chatAdmin.Feedback) {
	dto.ID = uint(feedbackPo.Id)
	dto.User = feedbackPo.User
	dto.MessageId = feedbackPo.MessageId
	dto.OpenKfId = feedbackPo.OpenKfId
	dto.Title = feedbackPo.Title
	dto.Content = feedbackPo.Content
	dto.Reply = feedbackPo.Reply
	dto.Status = feedbackPo.Status
	dto.CreatedAt = utils.TimeFormat(feedbackPo.CreatedAt)
	dto.UpdatedAt = utils.TimeFormat(feedbackPo.UpdatedAt)
	return

}
