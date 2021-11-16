package request

// User register structure
type Register struct {
	Username     string   `json:"userName"`
	Password     string   `json:"passWord"`
	NickName     string   `json:"nickName" gorm:"default:'QMPlusUser'"`
	HeaderImg    string   `json:"headerImg" gorm:"default:'https://qmplusimg.henrongyi.top/gva_header.jpg'"`
}

// User login structure
type Login struct {
	Username  string `json:"username"`  // 用户名
}

