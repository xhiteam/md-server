package core

import (
	"fmt"
	"github.com/kaijyin/md-server/server/global"
	"github.com/kaijyin/md-server/server/initialize"
	"time"

	"go.uber.org/zap"
)

type server interface {
	ListenAndServe() error
}

func RunWindowsServer() {
	// 初始化redis服务
	initialize.Redis()

	Router := initialize.Routers()

	Router.Static("/form-generator", "./resource/page")

	address := fmt.Sprintf(":%d", global.MD_CONFIG.System.Port)
	s := initServer(address, Router)
	// 保证文本顺序输出
	// In order to ensure that the text order output can be deleted
	time.Sleep(10 * time.Microsecond)
	global.MD_LOG.Info("server run success on ", zap.String("address", address))
	global.MD_LOG.Error(s.ListenAndServe().Error())
}
