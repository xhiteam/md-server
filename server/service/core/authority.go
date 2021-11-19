package core

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/kaijyin/md-server/server/global"
	"github.com/kaijyin/md-server/server/model/request"
	"github.com/kaijyin/md-server/server/model/response"
	"github.com/kaijyin/md-server/server/model/table"
	"github.com/kaijyin/md-server/server/utils"
	"gorm.io/gorm"
	"math"
	"strconv"
	"time"
)

const key = "20210604"

type AuthorityService struct {
}

var AuthorityServiceApp = new(AuthorityService)

type linkModel struct {
	ContextId uint
	Permission table.PermissionType
}
func (a *AuthorityService) CheckAuthority(userId uint, contextId uint, permission table.PermissionType) error {
	key := strconv.Itoa(int(userId)) + strconv.Itoa(int(contextId))
	db := global.MD_REDIS
	res, err := db.Get(ctx, key).Result()
	if err != nil {
		if !errors.Is(err, redis.Nil) {
			return err
		}
		//global.MD_LOG.Info("从磁盘中取",zap.Any("user",userId),zap.Any("user",contextId))
		control := table.Control{}
		if err = global.MD_DB.Where("user_id = ? and context_id = ?", userId, contextId).First(&control).Error; err != nil {
			return errors.New("权限不足")
		}
		timer := time.Duration(global.MD_CONFIG.JWT.ExpiresTime) * time.Second
		if control.Permission == table.Read {
			db.Set(ctx, key, "read",timer)
			res="read"
		} else {
			db.Set(ctx, key, "write",timer)
			res="write"
		}
	}else{
		//global.MD_LOG.Info("从内存中取",zap.Any("user",userId),zap.Any("user",contextId))
	}
	if res == "write" || permission == table.Read {
		return nil
	} else {
		return errors.New("权限不足")
	}
}
func (a *AuthorityService) GetAuthority(uid uint,fatherCatalogId uint,info response.ContextInfo, newPermission table.PermissionType,tx *gorm.DB)error  {
	err:=tx.Transaction(func(db *gorm.DB) error {
		info.Permission= newPermission
		err:=db.Create(table.Control{
			UserId:          uid,
			IsCatalog:       info.IsCatalog,
			ContextName:     info.ContextName,
			ContextId:       info.ContextId,
			FatherCatalogId: fatherCatalogId,
			Permission:      newPermission,
		}).Error;
		if err != nil {
			return err
		}
		key := strconv.Itoa(int(uid)) + strconv.Itoa(int(info.ContextId))
		timer := time.Duration(global.MD_CONFIG.JWT.ExpiresTime) * time.Second
		if newPermission== table.Read {
			global.MD_REDIS.Set(ctx, key, "read",timer)
		} else {
			global.MD_REDIS.Set(ctx, key, "write",timer)
		}
		if !info.IsCatalog{
			return nil
		}
		e1,infolist:=ContextServiceApp.GetContexts(request.GetContextsInfoReq{
			UID:             request.UID{UserId:uid},
			FatherCatalogId: info.ContextId,
			PageInfo:        request.PageInfo{
				Page:     1,
				PageSize: math.MaxInt,
			},
		})
		if e1 != nil {
			return err
		}
		for _, curinfo := range infolist.ContextsInfo {
			if e2:=a.GetAuthority(uid,info.ContextId,curinfo,newPermission,db);e2 != nil {
				return e2
			}
		}
		return nil
	})
	return err
}
func (a *AuthorityService) UpdateAuthority(uid uint,info response.ContextInfo, newPermission table.PermissionType,tx *gorm.DB) error{
	err:=tx.Transaction(func(db *gorm.DB) error {
		info.Permission= newPermission
		err:=db.Model(&table.Control{}).Where("user_id = ? and context_id = ?",uid,info.ContextId).Update("newPermission", newPermission).Error;
		if err != nil {
		  return err
		}
		if !info.IsCatalog{
			return nil
		}
		e1,infolist:=ContextServiceApp.GetContexts(request.GetContextsInfoReq{
			UID:             request.UID{UserId:uid},
			FatherCatalogId: info.ContextId,
			PageInfo:        request.PageInfo{
				Page:     1,
				PageSize: math.MaxInt,
			},
		})
		if e1 != nil {
			return err
		}
		key := strconv.Itoa(int(uid)) + strconv.Itoa(int(info.ContextId))
		timer := time.Duration(global.MD_CONFIG.JWT.ExpiresTime) * time.Second
		if newPermission== table.Read {
			global.MD_REDIS.Set(ctx, key, "read",timer)
		} else {
			global.MD_REDIS.Set(ctx, key, "write",timer)
		}
		for _, curinfo := range infolist.ContextsInfo {
			if e2:=a.UpdateAuthority(uid,curinfo,newPermission,db);e2 != nil {
				return e2
			}
		}
		return nil
	})
	return err
}
func (a *AuthorityService) CreateContextLink(req request.CreateContextLinkReq)(err error,contextLink response.ContextLinkResp) {
	var permission table.PermissionType
	if req.Permission=="read"{
		permission=table.Read
	}else if req.Permission=="write"{
		permission=table.Write
	}else {
		return  errors.New("权限控制字段(permission)错误"),contextLink
	}
	err=a.CheckAuthority(req.UserId,req.ContextId,permission)
	if err!=nil {
		return err,contextLink
	}
	var b []byte
    b,err=json.Marshal(linkModel{
		ContextId:  req.UserId,
		Permission: permission,
	})

	var link string
	link, err = utils.DesEncoding(string(b))
	if err!=nil{
		return err,contextLink
	}
	domainName:=global.MD_CONFIG.System.DomainName
	port:=fmt.Sprintf("%d",global.MD_CONFIG.System.Port)
    contextLink.ContextLink="https://"+domainName+":"+port+"/"+link
	return err,contextLink
}
func (a *AuthorityService) GetContextByLink(req request.GetContextByLinkReq )(err error,info response.ContextInfo){
	link:=linkModel{}
	err = json.Unmarshal([]byte(req.ContextLink), &link)
	if err!=nil{
		return err,info
	}

	controlRow:=table.Control{}
	if err=global.MD_DB.Where("context_id = ?",link.ContextId).First(&controlRow).Error;err!=nil{
	//确认是否存在,并得到文件元信息(名称,是否为目录)
		return err,info
	}
	//设置用户id
	controlRow.UserId=req.UserId
	e:=a.CheckAuthority(req.UserId,controlRow.FatherCatalogId,link.Permission)
	if e==nil{
		err= a.UpdateAuthority(req.UserId, response.ContextInfo{
			ContextId:   controlRow.ContextId,
			IsCatalog:   controlRow.IsCatalog,
			ContextName: controlRow.ContextName,
		},link.Permission,global.MD_DB)

	}else{
		err=a.GetAuthority(req.UserId,0,response.ContextInfo{
			ContextId:   controlRow.ContextId,
			IsCatalog:   controlRow.IsCatalog,
			ContextName: controlRow.ContextName,
		},link.Permission,global.MD_DB)
	}
	return err,info
}



