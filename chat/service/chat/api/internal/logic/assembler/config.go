package assembler

import (
	"chat/common/util"
	"chat/service/chat/api/internal/repository"
	"chat/service/chat/api/internal/types"
	"chat/service/chat/api/internal/vars"
	"chat/service/chat/model"
)

func POTODTOGetConfig(configPos []*model.Config) (dto *types.GetConfigResponse) {
	if len(configPos) <= 0 {
		return &types.GetConfigResponse{}
	}

	var info []types.Config
	for _, v := range configPos {
		info = append(info, POTODTOGetOne(v))
	}
	dto = &types.GetConfigResponse{
		List: info,
	}
	return
}

func POTODTOGetOne(configPos *model.Config) (dto types.Config) {
	value := configPos.Value
	if configPos.ConfigType == repository.ConfigTypeOpenAi {
		value = util.Desensitize(configPos.Value)
	}
	dto = types.Config{
		Id:          configPos.Id,
		ConfigType:  configPos.ConfigType,
		Name:        configPos.Name,
		Description: configPos.Description,
		Value:       value,
		CreatedAt:   configPos.CreatedAt.Format(vars.TimeFormat),
		UpdatedAt:   configPos.UpdatedAt.Format(vars.TimeFormat),
		CreatedBy:   configPos.CreatedBy,
		UpdatedBy:   configPos.UpdatedBy,
	}

	return
}
