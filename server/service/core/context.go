package core

import (
	"errors"
	"github.com/kaijyin/md-server/server/global"
	"github.com/kaijyin/md-server/server/model/request"
	"github.com/kaijyin/md-server/server/model/response"
	"github.com/kaijyin/md-server/server/model/table"
	"gorm.io/gorm"
	"math"
	"strconv"
)

type ContextService struct {
}

var ContextServiceApp = new(ContextService)

//先检查是否重命名,在权限控制表和文档表中添加

func (apiService *ContextService) CreateCatalog(req request.CreateCatalogReq) (err error, resp response.ContextInfo) {
	global.MD_DB.Transaction(func(db *gorm.DB) error {
		global.MD_LOG.Info("here")
		control := table.Control{}
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
		resp = response.ContextInfo{
			ContextId:   catalog.ID,
			CreatedAt:   control.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:   control.UpdatedAt.Format("2006-01-02 15:04:05"),
			IsCatalog:   true,
			ContextName: req.CatalogName,
		}
		key := strconv.Itoa(int(req.UserId)) + strconv.Itoa(int(catalog.ID))
		if err=RedisServiceApp.Set(key,"write");err != nil {
			return err
		}
		return nil
	})
	return err, resp
}
func (apiService *ContextService) CreateDocument(req request.CreateDocumentReq) (err error, resp response.ContextInfo) {
	err = global.MD_DB.Transaction(func(db *gorm.DB) error {
		control := table.Control{}
		err = db.Where("user_id = ? and father_catalog_id = ? and context_name = ?", req.UserId, req.FatherCatalogId, req.DocumentName).First(&control).Error
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			req.DocumentName = req.DocumentName + "(副本)"
		}
		document := table.Context{Content: req.Content}
		if err = db.Create(&document).Error; err != nil {
			return err
		}
		control = table.Control{
			UserId:          req.UserId,
			IsCatalog:       false,
			ContextName:     req.DocumentName,
			ContextId:       document.ID,
			FatherCatalogId: req.FatherCatalogId,
			Permission:      table.Write,
		}
		if err = db.Create(&control).Error; err != nil {
			return err
		}
		resp = response.ContextInfo{
			ContextId:   document.ID,
			CreatedAt:   control.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:   control.UpdatedAt.Format("2006-01-02 15:04:05"),
			IsCatalog:   true,
			ContextName: req.DocumentName,
		}
		key := strconv.Itoa(int(req.UserId)) + strconv.Itoa(int(document.ID))
		if err=RedisServiceApp.Set(key,"write");err != nil {
			return err
		}
		return nil
	})
	return err, resp
}

//删除文章
func (apiService *ContextService) DeleteDocument(req request.DeleteDocumentReq, tx *gorm.DB) (err error) {
	if err = AuthorityServiceApp.CheckAuthority(req.UserId, req.DocumentId, table.Write); err != nil {
		return err
	}
	err = tx.Transaction(func(db *gorm.DB) error {
		if err = db.Where("user_id = ? and context_id = ?", req.UserId, req.DocumentId).Delete(&table.Control{}).Error; err != nil {
			return err
		}
		if err = db.Delete(&table.Context{}, req.DocumentId).Error; err != nil {
			return err
		}
		key := strconv.Itoa(int(req.UserId)) + strconv.Itoa(int(req.DocumentId))
		if err=RedisServiceApp.Delete(key);err != nil {
			return err
		}
		return nil
	})
	return err
}

func (apiService *ContextService) DeleteCatalog(req request.DeleteCatalogReq, tx *gorm.DB) (err error) {
	if err = AuthorityServiceApp.CheckAuthority(req.UserId, req.CatalogId, table.Write); err != nil {
		return err
	}
	err = tx.Transaction(func(db *gorm.DB) error {
		if err = db.Where("context_id = ?", req.CatalogId).Delete(&table.Control{}).Error; err != nil {
			return err
		}
		sonReq := request.GetContextsInfoReq{
			UID:             req.UID,
			FatherCatalogId: req.CatalogId,
			PageInfo: request.PageInfo{
				Page:     1,
				PageSize: math.MinInt32,
			},
		}
		if err = db.Delete(&table.Context{}, req.CatalogId).Error; err != nil {
			return err
		}
		key := strconv.Itoa(int(req.UserId)) + strconv.Itoa(int(req.CatalogId))
		if err=RedisServiceApp.Delete(key);err != nil {
			return err
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
				}, db)
			} else {
				err = apiService.DeleteDocument(request.DeleteDocumentReq{
					UID:        req.UID,
					DocumentId: info.ContextId,
				}, db)
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
	db := global.MD_DB.Model(&table.Control{}).Select("context_id", "created_at", "updated_at", "is_catalog", "context_name").
		Where("user_id = ? and is_catalog = true and context_name like %?%", req.UserId, req.CatalogName).
		Offset(offset).Limit(limit).Order("created_at desc")
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
	db := global.MD_DB.Model(&table.Control{}).Select("context_id", "created_at", "updated_at",
		"is_catalog", "context_name", "permission").
		Where("user_id = ? and father_catalog_id = ?", req.UserId, req.FatherCatalogId).
		Offset(offset).Limit(limit).Order("created_at desc")
	if err = db.Scan(&list.ContextsInfo).Error; err != nil {
		return err, list
	}
	infos := list.ContextsInfo
	for i := range infos {
		infos[i].CreatedAt = global.NormalFormat(infos[i].CreatedAt)
		infos[i].UpdatedAt = global.NormalFormat(infos[i].UpdatedAt)
	}
	list.Total = len(list.ContextsInfo)
	return nil, list
}

//先检查用户有没有查看权限,有权限再获取

func (apiService *ContextService) GetContentById(req request.GetContentByIdReq) (err error, doc string) {
	if err = AuthorityServiceApp.CheckAuthority(req.UserId, req.DocumentId, table.Read); err != nil {
		return err, doc
	}
	if err = global.MD_DB.Select("content").Where("id = ?", req.DocumentId).Scan(&doc).Error; err != nil {
		return err, doc
	}
	return nil, doc
}

//先检查用户有没有权限修改,有权限再改

func (apiService *ContextService) UpdateDocumentContent(req request.UpdateContentReq) (err error) {
	if err = AuthorityServiceApp.CheckAuthority(req.UserId, req.DocumentId, table.Write); err != nil {
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
	if err = AuthorityServiceApp.CheckAuthority(req.UserId, req.ContextId, table.Write); err != nil {
		return err
	}
	db := global.MD_DB.Where("context_id = ?", req.ContextId)
	if err = db.Update("context_name", req.NewName).Error; err != nil {
		return err
	}
	return nil
}
