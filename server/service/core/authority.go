package core

import (
	"encoding/json"
	"errors"
	"github.com/kaijyin/md-server/server/global"
	"github.com/kaijyin/md-server/server/model/request"
	"github.com/kaijyin/md-server/server/model/response"
	"github.com/kaijyin/md-server/server/model/table"
	"github.com/kaijyin/md-server/server/utils"
	"gorm.io/gorm"
)

const key = "20210604"

type AuthorityService struct {
}

var AuthorityServiceApp = new(AuthorityService)

type linkModel struct {
	ContextId uint
	Permission table.PermissionType
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
	err=RedisServiceApp.CheckAuthority(req.UserId,req.ContextId,permission)
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
    contextLink.ContextLink="https://"+global.MD_CONFIG.System.DomainName+":"+global.MD_CONFIG.System.Port+"/"+link
	return err,contextLink
}

func (a *AuthorityService) GetContextByLink(req request.GetContextByLinkReq )(err error,info response.ContextInfo){
	link:=linkModel{}
	err = json.Unmarshal([]byte(req.ContextLink), &link)
	if err!=nil{
		return err,info
	}

	controlRow:=table.Control{}
	err=global.MD_DB.Where("context_id = ?",link.ContextId).First(&controlRow).Error
	//确认是否存在,并得到文件元信息(名称,是否为目录)
	if errors.Is(err, gorm.ErrRecordNotFound){
		return err,info
	}
	//设置用户id
	controlRow.UserId=req.UserId
	e:=RedisServiceApp.CheckAuthority(req.UserId,controlRow.FatherCatalogId,table.Read)
	if e==nil{
	  err=global.MD_DB.Model(&table.Control{}).Where("user_id = ? and context_id = ?",req.UserId,link.ContextId).
	  	Update("permission",link.Permission).Error
	  //TODO 如果是目录,后面的文件都要改权限
	}else{
		controlRow.FatherCatalogId=0
		controlRow.Permission=link.Permission
		err=global.MD_DB.Create(&controlRow).Error
		//TODO 如果是目录,那后面的也要添加权限
	}
	return err,info
}



