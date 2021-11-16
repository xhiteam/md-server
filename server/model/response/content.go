package response

type CreateContextResp struct {
	ContextId  uint `json:"contextId"`
}


type GetContextContentResp struct {
	Content string `json:"content"`
}


type ContextInfo struct {
	ContextId uint `json:"contextId"`
	IsCatalog bool `json:"isCatalog"`
	ContextName string `json:"contextName"`
}
type ContextInfoList struct {
	ContextsInfo []ContextInfo
	Total int `json:"total"`
}

