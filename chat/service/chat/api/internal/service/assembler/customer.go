package assembler

import (
	"chat/service/chat/api/internal/types"
	"chat/service/chat/api/internal/vars"
	"chat/service/chat/model"
	"strconv"
)

func POTODTOGetEvaluationSocket(pos []*model.ChatRecord) (dto []types.ChatRecordResponse) {
	if len(pos) <= 0 {
		return
	}

	for _, v := range pos {
		dto = append(dto, types.ChatRecordResponse{
			Id:               v.Id,
			RelationId:       strconv.Itoa(int(v.RelationId)),
			User:             v.User,
			OpenKfId:         v.OpenKfId,
			Content:          v.Content,
			ChatType:         v.ChatType,
			AnswerOrQuestion: v.AnswerOrQuestion,
			State:            v.State,
			StateId:          v.StateId,
			MessageType:      v.MessageType,
			CreatedAt:        v.CreatedAt.Format(vars.TimeFormat),
			UpdatedAt:        v.UpdatedAt.Format(vars.TimeFormat),
		})
	}
	return
}
