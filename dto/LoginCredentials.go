package dto

// 用户请求数据表单
type LoginCredentials struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}
