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
	"strconv"
)

var ConfigType int = 1

type ConfigService struct {
}

func (c *ConfigService) CreateConfig(Config *chatAdmin.Config) (err error) {
	Config.ConfigType = &ConfigType
	param := assembler.DTOTOPOAddConfig(Config)

	jsonParam, _ := json.Marshal(param)

	_, err = utils.Post(utils.GetChatServerHost()+vars.ChatAddConfigUri, jsonParam, nil)

	if err != nil {
		return
	}
	return
}

func (c *ConfigService) DeleteConfig(Config chatAdmin.Config) (err error) {
	param := assembler.DTOTOPODeleteConfig(Config)

	jsonParam, _ := json.Marshal(param)

	_, err = utils.Post(utils.GetChatServerHost()+vars.ChatDeleteConfigUri, jsonParam, nil)

	if err != nil {
		return
	}
	return
}

func (c *ConfigService) DeleteConfigByIds(ids request.IdsReq) (err error) {

	return
}

func (c *ConfigService) UpdateConfig(Config chatAdmin.Config) (err error) {
	Config.ConfigType = &ConfigType

	param := assembler.DTOTOPOUpdateConfig(&Config)

	jsonParam, _ := json.Marshal(param)

	_, err = utils.Post(utils.GetChatServerHost()+vars.ChatUpdateConfigUri, jsonParam, nil)

	if err != nil {
		return
	}
	return
}

func (c *ConfigService) GetConfig(id uint) (Config chatAdmin.Config, err error) {

	result, err := utils.Get(utils.GetChatServerHost()+vars.ChatFindConfigUri+"?id="+strconv.Itoa(int(id)), nil)
	if err != nil {
		return
	}
	var resultInfo clientStruct.Config
	err = json.Unmarshal(result, &resultInfo)
	if err != nil {
		return
	}

	if resultInfo.Id <= 0 {
		return
	}
	Config = assembler.POTODTOGetConfig(resultInfo)

	return
}

func (c *ConfigService) GetConfigInfoList(info chatAdminReq.ConfigOpenSearch) (list []chatAdmin.Config, total int64, err error) {
	var a int = 1
	info.ConfigType = &a
	param := map[string]string{
		"config_type": strconv.Itoa(*info.ConfigType),
	}

	result, err := utils.Get(utils.GetChatServerHost()+vars.ChatGetConfigUri+"?config_type="+strconv.Itoa(*info.ConfigType), param)

	if err != nil {
		return
	}
	var resultInfo clientStruct.GetConfigResponse
	err = json.Unmarshal(result, &resultInfo)
	if err != nil {
		return
	}
	if resultInfo.List == nil || len(resultInfo.List) <= 0 {
		return
	}
	list = assembler.POTODTOGetConfigList(resultInfo.List)

	return
}
