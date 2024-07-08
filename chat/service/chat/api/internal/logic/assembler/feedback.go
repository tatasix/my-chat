package assembler

import (
	"chat/common/util"
	"chat/service/chat/api/internal/types"
	"chat/service/chat/api/internal/vars"
	"chat/service/chat/model"
)

func POTODTOGetFeedbackList(feedbackPos []*model.Feedback) (feedbackResponse []types.FeedbackResponse) {
	if len(feedbackPos) <= 0 {
		return
	}
	for _, v := range feedbackPos {
		feedbackResponse = append(feedbackResponse, POTODTOFindFeedback(v))
	}

	return
}

func POTODTOFindFeedback(feedbackPo *model.Feedback) (feedbackResponse types.FeedbackResponse) {
	if feedbackPo == nil {
		return
	}
	feedbackResponse.Id = feedbackPo.Id
	feedbackResponse.User = feedbackPo.User
	feedbackResponse.MessageId = feedbackPo.MessageId
	feedbackResponse.OpenKfId = feedbackPo.OpenKfId
	feedbackResponse.Title = feedbackPo.Title
	feedbackResponse.Content = util.SqlToString(feedbackPo.Content)
	feedbackResponse.Reply = util.SqlToString(feedbackPo.Reply)
	feedbackResponse.Status = feedbackPo.Status
	feedbackResponse.CreatedAt = feedbackPo.CreatedAt.Format(vars.TimeFormat)
	feedbackResponse.UpdatedAt = feedbackPo.CreatedAt.Format(vars.TimeFormat)

	return
}
