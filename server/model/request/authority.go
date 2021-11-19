package request

// 名称模糊查询+分页+是否按照名称排序

type CreateContextLinkReq struct {
	UID
	ContextId uint `uri:"contextId"`
	Permission string `uri:"permission"`
}

type GetContextByLinkReq struct {
	UID
	ContextLink string `uri:"contextLink"`
}

