package table

import (
	"github.com/kaijyin/md-server/server/global"
)

type User struct {
	global.GVA_MODEL
	Username string `json:"userName" gorm:"index;unique;not null"`
}
