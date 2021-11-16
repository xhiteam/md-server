package service

import (
	"github.com/kaijyin/md-server/server/service/core"
	"github.com/kaijyin/md-server/server/service/system"
)

type ServiceGroup struct {
	System  system.ServiceGroup
	Core core.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
