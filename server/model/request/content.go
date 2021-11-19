package request

type UID struct {
	UserId uint `form:"userId"`
}
type CreateCatalogReq struct {
	UID
	FatherCatalogId uint   `uri:"fatherCatalogId"`
	CatalogName     string `uri:"catalogName"`
}

type CreateDocumentReq struct {
	UID
	FatherCatalogId uint   `uri:"fatherCatalogId"`
	DocumentName    string `uri:"documentName"`
	Content         string `uri:"content"`
}

type DeleteCatalogReq struct {
	UID
	CatalogId uint `uri:"catalogId"`
}
type DeleteDocumentReq struct {
	UID
	DocumentId uint `uri:"documentId"`
}

type GetContentByIdReq struct {
	UID
	DocumentId uint `uri:"documentId"`
}

type UpdateContextNameReq struct {
	UID
	ContextId uint   `uri:"contextId"`
	NewName   string `uri:"newName"`
}

type UpdateContentReq struct {
	UID
	DocumentId uint   `uri:"documentId"`
	NewContent string `uri:"newContent"`
}

// 名称模糊查询+分页+是否按照名称排序


type GetContextsInfoReq struct {
	UID
	FatherCatalogId uint `uri:"fatherCatalogId"`
	PageInfo
}

type GetCatalogsInfoByNameReq struct {
	UID
	CatalogName string `uri:"catalogName"`
	PageInfo
}
