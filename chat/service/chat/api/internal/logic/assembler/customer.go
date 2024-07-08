package assembler

import (
	"chat/service/chat/api/internal/types"
	"chat/service/chat/api/internal/vars"
	"chat/service/chat/model"
	"encoding/json"
)

func POTODTOGetCustomerList(customerPos []*model.CustomerConfig) (dto []types.CustomerConfig) {
	if len(customerPos) <= 0 {
		return
	}

	for _, v := range customerPos {
		dto = append(dto, POTODTOGetCustomer(v))
	}
	return
}

func POTODTOGetCustomer(customerPo *model.CustomerConfig) (dto types.CustomerConfig) {
	if customerPo == nil {
		return
	}
	var score float64
	if customerPo.Score.Valid {
		score = customerPo.Score.Float64
	}
	dto.Id = customerPo.Id
	dto.KfId = customerPo.KfId
	dto.KfName = customerPo.KfName
	dto.Prompt = customerPo.Prompt
	dto.PostModel = customerPo.PostModel
	dto.EmbeddingEnable = customerPo.EmbeddingEnable
	dto.EmbeddingMode = customerPo.EmbeddingMode
	dto.Score = score
	dto.TopK = customerPo.TopK
	dto.ClearContextTime = customerPo.ClearContextTime
	dto.CreatedAt = customerPo.CreatedAt.Format(vars.TimeFormat)
	dto.UpdatedAt = customerPo.UpdatedAt.Format(vars.TimeFormat)
	dto.Multiple = customerPo.Multiple
	dto.PromptStates = customerPo.PromptStates
	dto.Type = customerPo.Type
	//dto.Config = customerPo.Config
	dto.Summary = customerPo.Summary
	dto.Note = customerPo.Note
	dto.Description = customerPo.Description
	dto.PcImage = customerPo.PcImage
	dto.H5Image = customerPo.H5Image
	dto.Quote = customerPo.Quote
	dto.Quote2 = customerPo.Quote2
	var config model.CustomerConfigConfig
	_ = json.Unmarshal([]byte(customerPo.Config), &config)
	dto.Total = config.Total
	return
}

func POTODTOFindCustomerConfig(customerPo *model.CustomerConfig, promptPo []*model.Prompt) (dto types.FindCustomerConfigResult) {
	if customerPo == nil {
		return
	}
	var score float64
	if customerPo.Score.Valid {
		score = customerPo.Score.Float64
	}
	dto.Id = customerPo.Id
	dto.KfId = customerPo.KfId
	dto.KfName = customerPo.KfName
	dto.Prompt = customerPo.Prompt
	dto.PostModel = customerPo.PostModel
	dto.EmbeddingEnable = customerPo.EmbeddingEnable
	dto.EmbeddingMode = customerPo.EmbeddingMode
	dto.Score = score
	dto.TopK = customerPo.TopK
	dto.ClearContextTime = customerPo.ClearContextTime
	dto.CreatedAt = customerPo.CreatedAt.Format(vars.TimeFormat)
	dto.UpdatedAt = customerPo.UpdatedAt.Format(vars.TimeFormat)
	dto.Multiple = customerPo.Multiple
	dto.PromptStates = customerPo.PromptStates
	dto.Type = customerPo.Type
	//dto.Config = customerPo.Config
	dto.Summary = customerPo.Summary
	dto.Note = customerPo.Note
	dto.Description = customerPo.Description
	dto.PcImage = customerPo.PcImage
	dto.H5Image = customerPo.H5Image
	var config model.CustomerConfigConfig
	_ = json.Unmarshal([]byte(customerPo.Config), &config)
	dto.Total = config.Total

	if promptPo != nil && len(promptPo) > 0 {
		for _, v := range promptPo {
			dto.SubPrompt = append(dto.SubPrompt, types.SubPrompt{
				Id:        v.Id,
				StateId:   v.StateId,
				KfId:      v.KfId,
				Title:     v.Title,
				Prompt:    v.Prompt,
				UpdatedBy: v.UpdatedBy,
				CreatedBy: v.CreatedBy,
				CreatedAt: v.CreatedAt.Format(vars.TimeFormat),
				UpdatedAt: v.UpdatedAt.Format(vars.TimeFormat),
				IsDeleted: v.IsDeleted,
			})
		}
	}

	return
}

func POTODTOGetPayConfig(configPos []*model.CustomerConfig) (dto []types.GetPayInfoData) {
	if configPos == nil || len(configPos) <= 0 {
		return
	}

	for _, v := range configPos {
		dto = append(dto, types.GetPayInfoData{
			KfId:   v.KfId,
			KfName: v.KfName,
			Pay:    v.Pay,
		})
	}

	return
}
