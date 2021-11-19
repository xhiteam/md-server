package core

import (
	"errors"
	"github.com/kaijyin/md-server/server/global"
	"github.com/kaijyin/md-server/server/model/request"
	"github.com/kaijyin/md-server/server/model/response"
	"github.com/kaijyin/md-server/server/model/table"
	"gorm.io/gorm"
	"math"
)

type ContextService struct {
}

var ContextServiceApp = new(ContextService)

//先检查是否重命名,在权限控制表和文档表中添加

func (apiService *ContextService) CreateCatalog(req request.CreateCatalogReq) (err error, resp response.CreateContextResp) {
	global.MD_DB.Transaction(func(db *gorm.DB) error {
		control := table.Control{}
		err = db.Where("user_id = ? and father_catalog_id = ? and context_name = ?", req.UserId, req.FatherCatalogId, req.CatalogName).First(&control).Error
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("创建重名文件")
		}
		catalog := table.Context{}
		if err = db.Create(&catalog).Error; err != nil {
			return err
		}
		control = table.Control{
			UserId:          req.UserId,
			IsCatalog:       true,
			ContextName:     req.CatalogName,
			ContextId:       catalog.ID,
			FatherCatalogId: req.FatherCatalogId,
			Permission:      table.Write,
		}
		if err = db.Create(&control).Error; err != nil {
			return err
		}
		resp.ContextId = catalog.ID
		return nil
	})
	return err, resp
}
func (apiService *ContextService) CreateDocument(req request.CreateDocumentReq) (err error, resp response.CreateContextResp) {
	err = global.MD_DB.Transaction(func(db *gorm.DB) error {
		control := table.Control{}
		err = db.Where("user_id = ? and father_catalog_id = ? and context_name = ?", req.UserId, req.FatherCatalogId, req.DocumentName).First(&control).Error
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("创建重名文件")
		}
		catalog := table.Context{Content: req.Content}
		if err = db.Create(&catalog).Error; err != nil {
			return err
		}
		control = table.Control{
			UserId:          req.UserId,
			IsCatalog:       false,
			ContextName:     req.DocumentName,
			ContextId:       catalog.ID,
			FatherCatalogId: req.FatherCatalogId,
			Permission:      table.Write,
		}
		if err = db.Create(&control).Error; err != nil {
			return err
		}
		resp.ContextId = catalog.ID
		return nil
	})
	return err, resp
}

//删除文章
func (apiService *ContextService) DeleteDocument(req request.DeleteDocumentReq) (err error) {
	if err = RedisServiceApp.CheckAuthority(req.UserId, req.DocumentId, table.Write); err != nil {
		return err
	}
	err = global.MD_DB.Transaction(func(db *gorm.DB) error {
		if err = db.Delete(&table.Context{}, req.DocumentId).Error; err != nil {
			return err
		}
		if err = db.Where("user_id = ? and context_id = ?", req.UserId, req.DocumentId).Delete(&table.Control{}).Error; err != nil {
			return err
		}
		return nil
	})
	return err
}

func (apiService *ContextService) DeleteCatalog(req request.DeleteCatalogReq) (err error) {
	if err = RedisServiceApp.CheckAuthority(req.UserId, req.CatalogId, table.Write); err != nil {
		return err
	}
	err = global.MD_DB.Transaction(func(db *gorm.DB) error {
		if err = db.Delete(&table.Context{}, req.CatalogId).Error; err != nil {
			return err
		}
		if err = db.Where("context_id = ?", req.CatalogId).Delete(&table.Control{}).Error; err != nil {
			return err
		}
		sonReq := request.GetContextsInfoReq{
			UID:             req.UID,
			FatherCatalogId: req.CatalogId,
			PageList: request.PageList{
				PageInfo: request.PageInfo{
					Page:     1,
					PageSize: math.MinInt32,
				},
				Desc: false,
			},
		}
		resp := response.ContextInfoList{}
		if err, resp = apiService.GetContexts(sonReq); err != nil {
			return err
		}
		//递归删除
		for _, info := range resp.ContextsInfo {
			if info.IsCatalog {
				err = apiService.DeleteCatalog(request.DeleteCatalogReq{
					UID:       req.UID,
					CatalogId: info.ContextId,
				})
			} else {
				err = apiService.DeleteDocument(request.DeleteDocumentReq{
					UID:        req.UID,
					DocumentId: info.ContextId,
				})
			}
			if err != nil {
				return err
			}
		}
		return nil
	})
	return err
}

//直接搜索

func (apiService *ContextService) GetCatalogsByName(req request.GetCatalogsInfoByNameReq) (err error, resp response.ContextInfoList) {
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)
	db := global.MD_DB.Model(&table.Control{}).Select("context_id, is_catalog, context_name").
		Where("user_id = ? and is_catalog = true and context_name like %?%", req.UserId, req.CatalogName).
		Offset(offset).Limit(limit)
	if req.Desc {
		db = db.Order("context_name desc")
	} else {
		db = db.Order("context_name")
	}
	if err = db.Scan(&resp.ContextsInfo).Error; err != nil {
		return err, resp
	}
	resp.Total = len(resp.ContextsInfo)
	return nil, resp
}

//获取用户所有

func (apiService *ContextService) GetContexts(req request.GetContextsInfoReq) (err error, list response.ContextInfoList) {
	limit := req.PageSize
	offset := req.PageSize * (req.Page - 1)
	db := global.MD_DB.Model(&table.Control{}).Select("context_id", "is_catalog", "context_name").
		Where("user_id = ? and father_catalog_id = ?", req.UserId, req.FatherCatalogId).
		Offset(offset).Limit(limit)
	db = db.Order("is_catalog")
	if req.Desc {
		db = db.Order("context_name desc")
	} else {
		db = db.Order("context_name")
	}
	if err = db.Scan(&list.ContextsInfo).Error; err != nil {
		return err, list
	}
	list.Total = len(list.ContextsInfo)
	return nil, list
}

//先检查用户有没有查看权限,有权限再获取

func (apiService *ContextService) GetContentById(req request.GetContentByIdReq) (err error, content response.GetContextContentResp) {
	if err = RedisServiceApp.CheckAuthority(req.UserId, req.DocumentId, table.Read); err != nil {
		return err, content
	}
	if err = global.MD_DB.Select("content").Where("context_id = ?", req.DocumentId).Scan(&content.Content).Error; err != nil {
		return err, content
	}
	return nil, content
}

//先检查用户有没有权限修改,有权限再改

func (apiService *ContextService) UpdateDocumentContent(req request.UpdateContentReq) (err error) {
	if err = RedisServiceApp.CheckAuthority(req.UserId, req.DocumentId, table.Write); err != nil {
		return err
	}
	err = global.MD_DB.Model(table.Context{
		GVA_MODEL: global.GVA_MODEL{ID: req.DocumentId},
		Content:   "",
	}).Update("content", req.NewContent).Error
	if err != nil {
		return err
	}
	return nil
}
func (apiService *ContextService) UpdateContextName(req request.UpdateContextNameReq) (err error) {
	if err = RedisServiceApp.CheckAuthority(req.UserId, req.ContextId, table.Write); err != nil {
		return err
	}
	db:= global.MD_DB.Where("context_id = ?",req.ContextId)
	if err=db.Update("context_name",req.NewName).Error; err != nil {
		return err
	}
	return nil
}
