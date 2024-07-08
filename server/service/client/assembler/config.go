package assembler

import (
	"github.com/flipped-aurora/gin-vue-admin/server/model/chatAdmin"
	"github.com/flipped-aurora/gin-vue-admin/server/service/client/clientStruct"
	"github.com/flipped-aurora/gin-vue-admin/server/vars"
	"time"
)

func POTODTOGetConfigList(configPos []clientStruct.Config) (dto []chatAdmin.Config) {
	if len(configPos) <= 0 {
		return
	}

	for _, v := range configPos {
		dto = append(dto, POTODTOGetConfig(v))
	}
	return
}

func POTODTOGetConfig(configPo clientStruct.Config) (dto chatAdmin.Config) {

	configType := int(configPo.ConfigType)
	createdAt, _ := time.Parse(vars.TimeFormat, configPo.CreatedAt)
	updatedAt, _ := time.Parse(vars.TimeFormat, configPo.UpdatedAt)

	dto.ID = uint(configPo.Id)
	dto.ConfigType = &configType
	dto.Description = configPo.Description
	dto.Name = configPo.Name
	dto.Value = configPo.Value
	dto.CreatedAt = createdAt
	dto.UpdatedAt = updatedAt
	dto.CreatedBy = configPo.CreatedBy
	dto.UpdatedBy = configPo.UpdatedBy
	return
}

func DTOTOPOAddConfig(dto *chatAdmin.Config) (configPo clientStruct.AddConfigRequest) {

	configPo.ConfigType = int64(*dto.ConfigType)
	configPo.Description = dto.Description
	configPo.Value = dto.Value
	configPo.Name = dto.Name
	configPo.CreatedBy = dto.CreatedBy

	return
}

func DTOTOPOUpdateConfig(dto *chatAdmin.Config) (configPo clientStruct.UpdateConfigRequest) {

	configPo.Id = int64(dto.ID)
	configPo.ConfigType = int64(*dto.ConfigType)
	configPo.Description = dto.Description
	configPo.Value = dto.Value
	configPo.Name = dto.Name
	configPo.UpdateBy = dto.UpdatedBy

	return
}

func DTOTOPODeleteConfig(dto chatAdmin.Config) clientStruct.IdV2Request {
	return clientStruct.IdV2Request{Id: int64(dto.ID)}
}
