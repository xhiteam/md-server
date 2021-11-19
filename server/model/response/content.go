package response

import "github.com/kaijyin/md-server/server/model/table"

type ContextInfo struct {
	ContextId   uint                 `json:"contextId"`
	CreatedAt   string               `json:"createdAt"`
	UpdatedAt   string               `json:"updatedAt"`
	IsCatalog   bool                 `json:"isCatalog"`
	ContextName string               `json:"contextName"`
	Permission  table.PermissionType `json:"permission"`
}
type ContextInfoList struct {
	ContextsInfo []ContextInfo
	Total        int `json:"total"`
}
