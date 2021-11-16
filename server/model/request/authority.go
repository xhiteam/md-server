package request

// 名称模糊查询+分页+是否按照名称排序

type CreateContextLinkReq struct {
	UID
	ContextId uint `json:"contextId"`
	Permission string `json:"permission"`
}

type GetContextByLinkReq struct {
	UID
	ContextLink string `json:"contextLink"`
}

