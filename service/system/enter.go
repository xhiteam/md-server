package system

import "github.com/kaijyin/md-server/server/service/core"

type ServiceGroup struct {
	core.RedisService
	SystemConfigService
	UserService
}
