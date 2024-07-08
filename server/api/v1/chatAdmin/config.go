package chatAdmin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/model/chatAdmin"
	chatAdminReq "github.com/flipped-aurora/gin-vue-admin/server/model/chatAdmin/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/request"
	"github.com/flipped-aurora/gin-vue-admin/server/model/common/response"
	"github.com/flipped-aurora/gin-vue-admin/server/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ConfigApi struct {
}

var configService = service.ServiceGroupApp.ClientGroup.ConfigService

// CreateConfig 创建Config
// @Tags Config
// @Summary 创建Config
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body chatAdmin.Config true "创建Config"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /config/createConfig [post]
func (configApi *ConfigApi) CreateConfig(c *gin.Context) {
	var config chatAdmin.Config
	err := c.ShouldBindJSON(&config)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := configService.CreateConfig(&config); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Error(err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteConfig 删除Config
// @Tags Config
// @Summary 删除Config
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body chatAdmin.Config true "删除Config"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /config/deleteConfig [delete]
func (configApi *ConfigApi) DeleteConfig(c *gin.Context) {
	var config chatAdmin.Config
	err := c.ShouldBindJSON(&config)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := configService.DeleteConfig(config); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Error(err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteConfigByIds 批量删除Config
// @Tags Config
// @Summary 批量删除Config
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除Config"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /config/deleteConfigByIds [delete]
func (configApi *ConfigApi) DeleteConfigByIds(c *gin.Context) {
	var IDS request.IdsReq
	err := c.ShouldBindJSON(&IDS)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := configService.DeleteConfigByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Error(err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateConfig 更新Config
// @Tags Config
// @Summary 更新Config
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body chatAdmin.Config true "更新Config"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /config/updateConfig [put]
func (configApi *ConfigApi) UpdateConfig(c *gin.Context) {
	var config chatAdmin.Config
	err := c.ShouldBindJSON(&config)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := configService.UpdateConfig(config); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Error(err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindConfig 用id查询Config
// @Tags Config
// @Summary 用id查询Config
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query chatAdmin.Config true "用id查询Config"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /config/findConfig [get]
func (configApi *ConfigApi) FindConfig(c *gin.Context) {
	var config chatAdmin.Config
	err := c.ShouldBindQuery(&config)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if reconfig, err := configService.GetConfig(config.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Error(err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reconfig": reconfig}, c)
	}
}

// GetConfigList 分页获取Config列表
// @Tags Config
// @Summary 分页获取Config列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query chatAdminReq.ConfigSearch true "分页获取Config列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /config/getConfigList [get]
func (configApi *ConfigApi) GetConfigList(c *gin.Context) {
	var pageInfo chatAdminReq.ConfigOpenSearch
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if list, total, err := configService.GetConfigInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Error(err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
