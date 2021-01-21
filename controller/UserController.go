package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserInfo(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "ok",
	})
}
