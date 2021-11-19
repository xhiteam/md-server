package request

// Paging common input parameter structure

type PageInfo struct {
	Page     int `uri:"page" form:"page"`         // 页码
	PageSize int `uri:"pageSize" form:"pageSize"` // 每页大小
}


type Empty struct{}
