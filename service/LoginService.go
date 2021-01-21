package service

// 登录服务接口
type LoginService interface {
	LoginUser(username string, password string) bool
}

// 登录结构体
type loginInformation struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// 继承登录接口
func (login *loginInformation) LoginUser(username string, password string) bool {
	return login.Username == username && login.Password == password
}

// 实例化结构体
func NewLoginService() LoginService {
	return &loginInformation{
		Username: "lsrong",
		Password: "lsrong",
	}
}
