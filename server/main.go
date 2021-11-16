package main

import (
	"github.com/kaijyin/md-server/server/core"
	"github.com/kaijyin/md-server/server/global"
	"github.com/kaijyin/md-server/server/initialize"
)

//go:generate go env -w GO111MODULE=on
//go:generate go env -w GOPROXY=https://goproxy.cn,direct
//go:generate go mod tidy
//go:generate go mod download

// @title Swagger Example API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name x-token
// @BasePath /
func main() {
	global.MD_VP = core.Viper()      // 初始化Viper
	global.MD_LOG = core.Zap()       // 初始化zap日志库
	global.MD_DB = initialize.Gorm() // gorm连接数据库
	if global.MD_DB != nil {
		initialize.MysqlTables(global.MD_DB) // 初始化表
		// 程序结束前关闭数据库链接
		db, _ := global.MD_DB.DB()
		defer db.Close()
	}
	core.RunWindowsServer()
}
