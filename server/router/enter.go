package router

import (
	"github.com/kaijyin/md-server/server/router/core"
	"github.com/kaijyin/md-server/server/router/system"
)

type RouterGroup struct {
   Core core.RouterGroup
   System system.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
