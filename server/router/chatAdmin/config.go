package chatAdmin

import (
	"github.com/flipped-aurora/gin-vue-admin/server/api/v1"
	"github.com/flipped-aurora/gin-vue-admin/server/middleware"
	"github.com/gin-gonic/gin"
)

type ConfigOpenRouter struct {
}

// InitConfigOpenRouter 初始化 ConfigOpen 路由信息
func (s *ConfigOpenRouter) InitConfigOpenRouter(Router *gin.RouterGroup) {
	configRouter := Router.Group("config").Use(middleware.OperationRecord())
	configRouterWithoutRecord := Router.Group("config")
	var configApi = v1.ApiGroupApp.ChatAdminApiGroup.ConfigApi
	{
		configRouter.POST("createConfigOpen", configApi.CreateConfig)   // 新建ConfigOpen
		configRouter.DELETE("deleteConfigOpen", configApi.DeleteConfig) // 删除ConfigOpen
		configRouter.PUT("updateConfigOpen", configApi.UpdateConfig)    // 更新ConfigOpen
	}
	{
		configRouterWithoutRecord.GET("findConfigOpen", configApi.FindConfig)       // 根据ID获取ConfigOpen
		configRouterWithoutRecord.GET("getConfigOpenList", configApi.GetConfigList) // 获取ConfigOpen列表
	}
}
