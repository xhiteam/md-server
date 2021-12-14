package utils

var (
	CreateContextLinkVerify = Rules{"DocumentId": {NotEmpty()}, "Permission": {NotEmpty()}}
	GetContextByLinkVerify  = Rules{"ContextLink": {NotEmpty()}}

	CreateCatalogVerify  = Rules{ "CatalogName": {NotEmpty()}}
	CreateDocumentVerify = Rules{ "DocumentName": {NotEmpty()},
		"Content": {NotEmpty()}}

	CatalogIdVerify  = Rules{"CatalogId": {NotEmpty()}}
	DocumentIdVerify = Rules{"DocumentId": {NotEmpty()}}

	UpdateContextNameVerify     = Rules{"ContextId": {NotEmpty()}, "NewName": {NotEmpty()}}
	UpdateDocumentContentVerify = Rules{"DocumentId": {NotEmpty()}, "NewContent": {NotEmpty()}}

	GetCatalogsInfoByNameVerify = Rules{"CatalogName": {NotEmpty()}, "Page": {NotEmpty()}, "PageSize": {NotEmpty()}}
	GetContextsInfoVerify       = Rules{"UID": {NotEmpty()},"Page": {NotEmpty()}, "PageSize": {NotEmpty()}}

	LoginVerify          = Rules{"Username": {NotEmpty()}}
	UserIdVerify          = Rules{"UID": {NotEmpty()}}
)
