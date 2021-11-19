package system

import (
	"github.com/kaijyin/md-server/server/global"
	"github.com/kaijyin/md-server/server/model/request"
	"github.com/kaijyin/md-server/server/model/response"
	"github.com/kaijyin/md-server/server/model/table"
	"github.com/kaijyin/md-server/server/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type BaseApi struct {
}

// @Tags Base
// @Summary 用户登录
// @Produce  application/json
// @Param data body systemReq.Login true "用户名, 密码, 验证码"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"登陆成功"}"
// @Router /base/login [post]
func (b *BaseApi) Login(c *gin.Context) {
	var l request.Login
	_ = c.ShouldBind(&l)
	if err := utils.Verify(l, utils.LoginVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	u := &table.User{Username: l.Username}
	if err, user := userService.Login(u); err != nil {
		global.MD_LOG.Error("登录失败,未知错误!", zap.Any("err", err))
		response.FailWithMessage("登录失败,未知错误", c)
	} else {
		b.tokenNext(c, *user)
	}
}

// 登录以后签发jwt
func (b *BaseApi) tokenNext(c *gin.Context, user table.User) {
	j := &utils.JWT{SigningKey: []byte(global.MD_CONFIG.JWT.SigningKey)} // 唯一签名
	claims := j.CreateClaims(request.BaseClaims{
		ID: user.ID,
	})

	global.MD_LOG.Info("login",zap.Any("userid",user.ID))
	token, err := j.CreateToken(claims)
	if err != nil {
		global.MD_LOG.Error("获取token失败!", zap.Any("err", err))
		response.FailWithMessage("获取token失败", c)
		return
	}
	response.OkWithDetailed(response.LoginResponse{
		User:      user.Username,
		Token:     token,
		ExpiresAt: claims.StandardClaims.ExpiresAt * 1000,
	}, "登录成功", c)
	return
}



// @Tags User
// @Summary 删除用户
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.GetById true "用户ID"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /user/deleteUser [delete]
func (b *BaseApi) DeleteUser(c *gin.Context) {
	userId := utils.GetUserID(c)
	uid := request.UID{UserId: userId}
	if err := userService.DeleteUser(uid); err != nil {
		global.MD_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}
