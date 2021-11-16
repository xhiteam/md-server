package system

import (
	"github.com/kaijyin/md-server/server/global"
	"github.com/kaijyin/md-server/server/model/request"
	"github.com/kaijyin/md-server/server/model/table"
)



type UserService struct {
}


func (userService *UserService) Login(u *table.User) (err error, userInter *table.User) {
	var user table.User
	global.MD_DB.Create(&u)
	return err, &user
}

func (userService *UserService) DeleteUser(req request.UID) (err error) {
	err = global.MD_DB.Delete(&table.User{},req.UserId).Error
	return err
}
