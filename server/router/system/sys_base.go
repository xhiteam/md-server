package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/kaijyin/md-server/server/api/v1"
)

type BaseRouter struct {
}

func (s *BaseRouter) InitBaseRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	var baseApi = v1.ApiGroupApp.SystemApiGroup.BaseApi
	{
		Router.POST("login", baseApi.Login)
	}
	return Router
}
