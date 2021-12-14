package core

import (
	"github.com/gin-gonic/gin"
	"github.com/kaijyin/md-server/server/global"
	"github.com/kaijyin/md-server/server/model/request"
	"github.com/kaijyin/md-server/server/model/response"
	"github.com/kaijyin/md-server/server/utils"
	"go.uber.org/zap"
)

type AuthorityApi struct {
}


func (a *AuthorityApi) CreateContextLink(c *gin.Context) {
	var req request.CreateContextLinkReq
	_ = c.ShouldBindUri(&req)
	uid,_:=c.Get("userId")
	req.UserId=uid.(uint)
	if err := utils.Verify(req, utils.CreateContextLinkVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, contextLink := AuthorityService.CreateContextLink(req); err != nil {
		global.MD_LOG.Error("分享链接创建失败!", zap.Any("err", err))
		response.FailWithMessage("分享链接创建失败"+err.Error(), c)
	} else {
		response.OkWithData(contextLink,c)
	}
}


func (a *AuthorityApi) GetContextByLink(c *gin.Context) {
	var req request.GetContextByLinkReq
	_ = c.ShouldBindUri(&req)
	uid,_:=c.Get("userId")
	req.UserId=uid.(uint)
	if err := utils.Verify(req, utils.GetContextByLinkVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err,contextInfo:= AuthorityService.GetContextByLink(req); err != nil {
		global.MD_LOG.Error("共享文件获取失败!", zap.Any("err", err))
		response.FailWithMessage("共享文件获取失败"+err.Error(), c)
	} else {
		response.OkWithDetailed(contextInfo, "共享文件获取成功", c)
	}
}
