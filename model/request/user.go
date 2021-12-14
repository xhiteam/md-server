package request


// User login structure
type Login struct {
	Username  string  `form:"username" json:"username"`  // 用户名
}

