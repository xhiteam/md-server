package response

import (
	"github.com/kaijyin/md-server/server/model/table"
)

type SysUserResponse struct {
	User table.User `json:"user"`
}

type LoginResponse struct {
	User      string `json:"user"`
	Token     string `json:"token"`
	ExpiresAt int64  `json:"expiresAt"`
}
