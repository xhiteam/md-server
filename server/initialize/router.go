package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/kaijyin/md-server/server/global"
	"github.com/kaijyin/md-server/server/middleware"
	"github.com/kaijyin/md-server/server/router"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// 初始化总路由

func Routers() *gin.Engine {
	var Router = gin.Default()

	//https
	//Router.Use(middleware.LoadTls())
	//global.MD_LOG.Info("use middleware tls")

	// 打开跨域请求
	Router.Use(middleware.Cors())
	global.MD_LOG.Info("use middleware cors")

	//swagger
	Router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	global.MD_LOG.Info("register swagger handler")


	PublicGroup := Router.Group("")
	{
		// 健康监测
		PublicGroup.GET("/health", func(c *gin.Context) {
			c.JSON(200, "ok")
		})
	}
	//获取路由组实例
	systemRouter := router.RouterGroupApp.System
	{
		systemRouter.InitBaseRouter(PublicGroup) // 注册基础功能路由,不做鉴权
	}
	PrivateGroup := Router.Group("")
	PrivateGroup.Use(middleware.JWTAuth())  //token鉴权
	{
		systemRouter.InitUserRouter(PrivateGroup)                   // 注册用户路由
		systemRouter.InitSystemRouter(PrivateGroup)                 // system相关路由
	}
	coreRouter:=router.RouterGroupApp.Core
    {
    	coreRouter.InitContextRouter(PrivateGroup)             //文章操作相关路由
    	coreRouter.InitAuthorityRouter(PrivateGroup)           //文章权限控制相关路由
	}
	global.MD_LOG.Info("router register success")
	return Router
}
