package assembler

import (
	"chat/service/chat/api/internal/types"
	"chat/service/chat/api/internal/vars"
	"chat/service/chat/model"
)

func POTODTOGetApplicationList(applicationPos []*model.ApplicationConfig) (dto []types.ApplicationConfig) {
	if len(applicationPos) <= 0 {
		return
	}

	for _, v := range applicationPos {
		dto = append(dto, POTODTOGetApplication(v))
	}
	return
}

func POTODTOGetApplication(applicationPos *model.ApplicationConfig) (dto types.ApplicationConfig) {
	if applicationPos == nil {
		return
	}
	var score float64
	if applicationPos.Score.Valid {
		score = applicationPos.Score.Float64
	}
	dto.AgentId = int(applicationPos.AgentId)
	dto.Id = applicationPos.Id
	dto.AgentSecret = applicationPos.AgentSecret
	dto.AgentName = applicationPos.AgentName
	dto.Model = applicationPos.Model
	dto.PostModel = applicationPos.PostModel
	dto.BasePrompt = applicationPos.BasePrompt
	dto.Welcome = applicationPos.Welcome
	dto.GroupEnable = applicationPos.GroupEnable
	dto.EmbeddingEnable = applicationPos.EmbeddingEnable
	dto.EmbeddingMode = applicationPos.EmbeddingMode
	dto.Score = score
	dto.TopK = int(applicationPos.TopK)
	dto.ClearContextTime = int(applicationPos.ClearContextTime)
	dto.GroupName = applicationPos.GroupName
	dto.CreatedAt = applicationPos.CreatedAt.Format(vars.TimeFormat)
	dto.UpdatedAt = applicationPos.UpdatedAt.Format(vars.TimeFormat)

	return
}
