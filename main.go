package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jwt-demo/controller"
	"github.com/jwt-demo/middleware"
	"github.com/jwt-demo/service"
)

func main() {
	var (
		loginService    = service.NewLoginService()
		jwtService      = service.NewJWAuthService()
		loginController = controller.LoginHandler(loginService, jwtService)
	)

	server := gin.Default()

	// 登录认证
	server.POST("/login", loginController.Login)

	// 鉴权
	apiGroups := server.Group("/api")
	apiGroups.Use(middleware.JWTAuthorize())
	{
		apiGroups.GET("/user/info", controller.UserInfo)
	}
	address := ":8888"

	err := server.Run(address)
	if err != nil {
		fmt.Printf("Failed to run server:%v ", err.Error())
	}
}
