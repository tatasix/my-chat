package assembler

import (
	"chat/service/chat/api/internal/types"
	"chat/service/chat/api/internal/vars"
	"chat/service/chat/model"
	"strconv"
)

func POTODTOGetChatList(chatPos, answer []*model.ChatRecord, userPos []*model.WechatUser, applicationPos []*model.ApplicationConfig, customerPos []*model.CustomerConfig) (dto []types.ChatResponse) {
	if len(chatPos) <= 0 {
		return
	}
	userMap := make(map[string]string)
	customerMap := make(map[string]string)
	applicationMap := make(map[int64]string)
	answerMap := make(map[int64]*model.ChatRecord)
	if len(userPos) > 0 {
		for _, user := range userPos {
			userMap[user.User] = user.Nickname
		}
	}

	if len(applicationPos) > 0 {
		for _, application := range applicationPos {
			applicationMap[application.AgentId] = application.AgentName
		}
	}

	if len(customerPos) > 0 {
		for _, customer := range customerPos {
			customerMap[customer.KfId] = customer.KfName
		}
	}

	if len(answer) > 0 {
		for _, answerSon := range answer {
			answerMap[answerSon.RelationId] = answerSon
		}
	}

	for _, v := range chatPos {
		u := v.User
		var a, c, answer string

		if value1, ok1 := userMap[u]; ok1 {
			u = value1
		}
		if value2, ok2 := applicationMap[v.AgentId]; ok2 {
			a = value2
		}
		if value3, ok3 := customerMap[v.OpenKfId]; ok3 {
			c = value3
		}
		if value4, ok4 := answerMap[v.Id]; ok4 {
			answer = value4.Content
		}
		dto = append(dto, types.ChatResponse{
			Id:             v.Id,
			User:           u,
			ChatRecordUser: v.User,
			MessageId:      v.MessageId,
			OpenKfId:       c,
			AgentId:        a,
			ReqContent:     v.Content,
			ResContent:     answer,
			CreatedAt:      v.CreatedAt.Format(vars.TimeFormat),
			UpdatedAt:      v.UpdatedAt.Format(vars.TimeFormat),
		})
	}
	return
}

func POTODTOGetChatRecordList(chatRecordPos []*model.ChatRecord) (dto []types.ChatRecordResponse) {
	if len(chatRecordPos) <= 0 {
		return
	}

	for _, v := range chatRecordPos {

		dto = append(dto, types.ChatRecordResponse{
			ChatRecordId:     strconv.FormatInt(v.Id, 10),
			Id:               v.Id,
			RelationId:       strconv.Itoa(int(v.RelationId)),
			User:             v.User,
			UserName:         "",
			MessageId:        v.MessageId,
			OpenKfId:         v.OpenKfId,
			OpenKfName:       "",
			AgentId:          v.AgentId,
			Content:          v.Content,
			Emoji:            v.Emoji,
			ChatType:         v.ChatType,
			AnswerOrQuestion: v.AnswerOrQuestion,
			MessageType:      v.MessageType,
			CreatedAt:        v.CreatedAt.Format(vars.TimeFormat),
			UpdatedAt:        v.UpdatedAt.Format(vars.TimeFormat),
		})
	}
	return
}

func POTODTOGetChatExportList(chatPos []*model.ChatRecord, userPos []*model.WechatUser, applicationPos []*model.ApplicationConfig, customerPos []*model.CustomerConfig) (dto []*model.ExportChatRecord) {
	if len(chatPos) <= 0 {
		return
	}
	userMap := make(map[string]string)
	customerMap := make(map[string]string)
	applicationMap := make(map[int64]string)
	if len(userPos) > 0 {
		for _, user := range userPos {
			userMap[user.User] = user.Nickname
		}
	}

	if len(applicationPos) > 0 {
		for _, application := range applicationPos {
			applicationMap[application.AgentId] = application.AgentName
		}
	}

	if len(customerPos) > 0 {
		for _, customer := range customerPos {
			customerMap[customer.KfId] = customer.KfName
		}
	}

	for _, v := range chatPos {
		u := v.User
		var a, c string

		if value1, ok1 := userMap[u]; ok1 {
			u = value1
		}
		if value2, ok2 := applicationMap[v.AgentId]; ok2 {
			a = value2
		}
		if value3, ok3 := customerMap[v.OpenKfId]; ok3 {
			c = value3
		}

		dto = append(dto, &model.ExportChatRecord{
			Id:               v.Id,
			RelationId:       v.RelationId,
			UserName:         u,
			MessageId:        v.MessageId,
			OpenKfName:       c,
			AgentName:        a,
			Content:          v.Content,
			Emoji:            v.Emoji,
			ChatType:         v.ChatType,
			AnswerOrQuestion: v.AnswerOrQuestion,
			MessageType:      v.MessageType,
			CreatedAt:        v.CreatedAt.Format(vars.TimeFormat),
			UpdatedAt:        v.UpdatedAt.Format(vars.TimeFormat),
		})
	}
	return
}
