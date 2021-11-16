package v1

import (
	"github.com/kaijyin/md-server/server/api/v1/core"
	"github.com/kaijyin/md-server/server/api/v1/system"
)

type ApiGroup struct {
	SystemApiGroup system.SysGroup
	CoreApiGroup   core.CoreGroup
}

var ApiGroupApp = new(ApiGroup)
