package system

import (

	"github.com/gin-gonic/gin"
	v1 "github.com/kaijyin/md-server/server/api/v1"
	"github.com/kaijyin/md-server/server/middleware"
)

type SysRouter struct {
}

func (s *SysRouter) InitSystemRouter(Router *gin.RouterGroup) {
	sysRouter := Router.Group("system").Use(middleware.OperationRecord())
	var systemApi = v1.ApiGroupApp.SystemApiGroup.SystemApi
	{
		sysRouter.GET("config", systemApi.GetSystemConfig) // 获取配置文件内容
		sysRouter.PUT("config", systemApi.SetSystemConfig) // 设置配置文件内容
	}
	{
		sysRouter.GET("info", systemApi.GetServerInfo)     // 获取服务器信息
		sysRouter.POST("reload", systemApi.ReloadSystem)       // 重启服务
	}
}
