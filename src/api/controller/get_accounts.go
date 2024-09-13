package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAccount(ctx *gin.Context) {
	user_name, exists := ctx.Get("user")
	if !exists {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"user_name": user_name,
	})
}
