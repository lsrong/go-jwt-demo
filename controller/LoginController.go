package controller

import (
	"net/http"

	"github.com/jwt-demo/dto"
	"github.com/jwt-demo/service"

	"github.com/gin-gonic/gin"
)

type LoginController interface {
	Login(ctx *gin.Context)
}
type loginController struct {
	loginService service.LoginService
	jwtService   service.JWTService
}

func LoginHandler(loginService service.LoginService, jwtService service.JWTService) LoginController {
	return &loginController{
		loginService: loginService,
		jwtService:   jwtService,
	}
}

// 登录业务
func (controller *loginController) Login(ctx *gin.Context) {
	var credential dto.LoginCredentials

	err := ctx.ShouldBind(&credential)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    -1,
			"message": "用户名密码参数必填",
		})
		return
	}
	isUser := controller.loginService.LoginUser(credential.Username, credential.Password)
	if !isUser {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "用户名或者密码不正确",
		})
		return
	}
	token := controller.jwtService.CreateToken(credential.Username)
	ctx.JSON(http.StatusOK, gin.H{
		"code":  0,
		"token": token,
	})
}
