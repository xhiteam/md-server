package system

import (
	"github.com/kaijyin/md-server/server/service"
)

type SysGroup struct {
	BaseApi
	SystemApi
}

var jwtService = service.ServiceGroupApp.System.RedisService
var userService = service.ServiceGroupApp.System.UserService
var systemConfigService = service.ServiceGroupApp.System.SystemConfigService
