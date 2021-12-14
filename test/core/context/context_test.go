package context

import (
	"github.com/kaijyin/md-server/server/model/request"
	"github.com/kaijyin/md-server/server/service/core"
	"testing"
)

//先检查是否重命名,在权限控制表和文档表中添加

func TestCreateCatalog(t *testing.T) {
	err, _ := core.ContextServiceApp.CreateCatalog(request.CreateCatalogReq{
		UID:             request.UID{UserId: 1},
		FatherCatalogId: 0,
		CatalogName:     "目录1",
	})
	if err != nil {
		t.Error(err)
	}
	t.Log("创建目录成功")
}
func  TestCreateDocument(t *testing.T) {

}

//删除文章
func TestDeleteDocument(t *testing.T)  {

}

func DeleteCatalog(t *testing.T)  {

}

//直接搜索

func TestGetCatalogsByName(t *testing.T)  {

}

//获取用户所有

func TestGetContexts(t *testing.T)  {

}

//先检查用户有没有查看权限,有权限再获取

func TestGetContentById(t *testing.T) {

}

//先检查用户有没有权限修改,有权限再改

func TestUpdateDocumentContent(t *testing.T){
}
func TestUpdateContextName(t *testing.T) {

}
