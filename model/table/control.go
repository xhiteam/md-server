package table

import (
	"github.com/kaijyin/md-server/server/global"
)

const (
	Read  = 1
	Write = 2
)

type PermissionType int

type Control struct {
	global.GVA_MODEL
	UserId          uint           `json:"userId" gorm:"not null"`                    // 角色ID
	IsCatalog       bool           `json:"isCatalog" gorm:"not null;default true"`    //标志位,是否为目录
	ContextName     string         `json:"contextName" gorm:"not null"`               // 目录/文章名称
	ContextId       uint           `json:"contextId" gorm:"not null"`                 //目录/文章Id
	FatherCatalogId uint           `json:"fatherContextId" gorm:"not null;default 0"` // 父目录Id,用于查找和递归删除
	Permission      PermissionType `json:"core" gorm:"not null;"`                     // 用户对文档所有权限
	ext             string         `json:"ext"`                                       // 扩展字段
	User            User           // 外键引用User表
	Context         Context
}

//```
//json
//{
//"UID":"所属用户Id",
//"contextId":"文档/目录id",
//"level":"所在层级",
//"flag":"目录文档标志位",
//"menuname":"文档/目录名称",
//"fatherId":"int",//父目录id,根目录为0,用于递归删除
//"permision":"权限int",//三种权限,owner/1,read/2,write/3
//"ext":"扩展字段",
//}
//```
