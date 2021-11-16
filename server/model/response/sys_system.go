package response

import "github.com/kaijyin/md-server/server/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}

type System struct {
	Config config.Server `json:"config"`
}