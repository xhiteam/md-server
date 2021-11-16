package core

import (
	"github.com/kaijyin/md-server/server/global"
	"github.com/kaijyin/md-server/server/model/request"
	"github.com/kaijyin/md-server/server/model/response"
	"github.com/kaijyin/md-server/server/utils"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type ContextApi struct {
}


func (s *ContextApi) CreateCatalog(c *gin.Context) {
	var req request.CreateCatalogReq
	_ = c.ShouldBindJSON(&req)
	if err := utils.Verify(req, utils.CreateCatalogVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err,resp := ContextService.CreateCatalog(req); err != nil {
		global.MD_LOG.Error("创建目录失败!", zap.Any("err", err))
		response.FailWithMessage("创建目录失败", c)
	} else {
		response.OkWithDetailed(resp,"创建目录成功", c)
	}
}

func (s *ContextApi) CreateDocument(c *gin.Context) {
	var req request.CreateDocumentReq
	_ = c.ShouldBindJSON(&req)
	if err := utils.Verify(req, utils.CreateDocumentVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err,resp:= ContextService.CreateDocument(req); err != nil {
		global.MD_LOG.Error("创建文档失败!", zap.Any("err", err))
		response.FailWithMessage("创建文档失败", c)
	} else {
		response.OkWithDetailed(resp,"创建文档成功", c)
	}
}

func (s *ContextApi) DeleteCatalog(c *gin.Context) {
	var req request.DeleteCatalogReq
	_ = c.ShouldBindJSON(&req)
	if err := utils.Verify(req, utils.CatalogIdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := ContextService.DeleteCatalog(req); err != nil {
		global.MD_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func (s *ContextApi) DeleteDocument(c *gin.Context) {
	var req request.DeleteDocumentReq
	_ = c.ShouldBindJSON(&req)
	if err := utils.Verify(req, utils.DocumentIdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := ContextService.DeleteDocument(req); err != nil {
		global.MD_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

func (s *ContextApi) UpdateContextName(c *gin.Context) {
	var req request.UpdateContextNameReq
	_ = c.ShouldBindJSON(&req)
	if err := utils.Verify(req, utils.UpdateContextNameVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := ContextService.UpdateContextName(req); err != nil {
		global.MD_LOG.Error("修改失败!", zap.Any("err", err))
		response.FailWithMessage("修改失败", c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}

func (s *ContextApi) UpdateDocumentContent(c *gin.Context) {
	var req request.UpdateContentReq
	_ = c.ShouldBindJSON(&req)
	if err := utils.Verify(req, utils.UpdateDocumentContentVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err := ContextService.UpdateDocumentContent(req); err != nil {
		global.MD_LOG.Error("修改失败!", zap.Any("err", err))
		response.FailWithMessage("修改失败", c)
	} else {
		response.OkWithMessage("修改成功", c)
	}
}




func (s *ContextApi) GetCatalogsInfoByName(c *gin.Context) {
	var req request.GetCatalogsInfoByNameReq
	_ = c.ShouldBindJSON(&req)
	if err := utils.Verify(req.PageInfo, utils.GetCatalogsInfoByNameVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, resp:= ContextService.GetCatalogsByName(req); err != nil {
		global.MD_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(resp, "获取成功", c)
	}
}

func (s *ContextApi) GetContextsInfo(c *gin.Context) {
	var req request.GetContextsInfoReq
	_ = c.ShouldBindJSON(&req)
	if err := utils.Verify(req, utils.GetContextsInfoVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if err, resp := ContextService.GetContexts(req); err != nil {
		global.MD_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(resp, "获取成功", c)
	}
}

func (s *ContextApi) GetContentById(c *gin.Context) {
	var req request.GetContentByIdReq
	_ = c.ShouldBindJSON(&req)
	if err := utils.Verify(req, utils.DocumentIdVerify); err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	if	err, context := ContextService.GetContentById(req);err != nil {
		global.MD_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithData(context, c)
	}
}