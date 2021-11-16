package core

import (
	"github.com/kaijyin/md-server/server/service"
)

type CoreGroup struct {
	AuthorityApi
	ContextApi
}
var ContextService = service.ServiceGroupApp.Core.ContextService
var AuthorityService = service.ServiceGroupApp.Core.AuthorityService
var jwtService = service.ServiceGroupApp.System.RedisService

