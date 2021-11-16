package request

type UID struct {
	UserId uint `json:"userId"`
}
type CreateCatalogReq struct {
	UID
	FatherCatalogId uint   `json:"fatherCatalogId"`
	CatalogName     string `json:"catalogName"`
}

type CreateDocumentReq struct {
	UID
	FatherCatalogId uint   `json:"fatherCatalogId"`
	DocumentName    string `json:"documentName"`
	Content         string `json:"content"`
}

type DeleteCatalogReq struct {
	UID
	CatalogId uint `json:"catalogId"`
}
type DeleteDocumentReq struct {
	UID
	DocumentId uint `json:"documentId"`
}

type GetContentByIdReq struct {
	UID
	DocumentId uint `json:"documentId"`
}

type UpdateContextNameReq struct {
	UID
	ContextId uint   `json:"contextId"`
	NewName   string `json:"newName"`
}

type UpdateContentReq struct {
	UID
	DocumentId uint   `json:"documentId"`
	NewContent string `json:"newContent"`
}

// 名称模糊查询+分页+是否按照名称排序

type PageList struct {
	PageInfo
	Desc bool `json:"desc"` // 排序方式:升序false(默认)|降序true
}

type GetContextsInfoReq struct {
	UID
	FatherCatalogId uint `json:"fatherCatalogId"`
	PageList
}

type GetCatalogsInfoByNameReq struct {
	UID
	CatalogName string `json:"catalogName"`
	PageList
}
