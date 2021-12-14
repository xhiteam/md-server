package table

import (
	"github.com/kaijyin/md-server/server/global"
)


type Context struct {
	global.GVA_MODEL
	Content     string             `json:"content" gorm:"type:text"`                     // 文章内容
}