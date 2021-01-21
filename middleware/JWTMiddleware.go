package middleware

import (
	"fmt"
	"net/http"

	"github.com/jwt-demo/service"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

const BearerSchema = "Bearer "

func JWTAuthorize() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorization := ctx.GetHeader("Authorization")
		if len(authorization) <= 0 {
			fmt.Println("Missing authorization header")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    -1,
				"message": "Unauthorized",
			})
			return
		}
		token := authorization[len(BearerSchema):]

		jwtToken, err := service.NewJWAuthService().ValidateToken(token)
		if err != nil || !jwtToken.Valid {
			fmt.Printf("Unauthrized token,error:%v \n", err)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"code":    -1,
				"message": "Unauthorized",
			})
			return
		}

		claims := jwtToken.Claims.(jwt.MapClaims)
		ctx.Set("claims", claims)
		ctx.Next()
	}
}
