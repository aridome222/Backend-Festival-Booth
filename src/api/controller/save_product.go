package controller

import (
	"net/http"

	"github.com/aridome222/Backend-Festival-Booth/usecase"
	"github.com/gin-gonic/gin"
)

type SaveProductController struct {
	uc usecase.SaveProductUseCase
}

func NewSaveProductController(uc usecase.SaveProductUseCase) SaveProductController {
	return SaveProductController{
		uc: uc,
	}
}

func (con SaveProductController) SaveProduct(ctx *gin.Context) {
	var input usecase.SaveProductInputDTO

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// cookieからユーザー名を取得
	user_name, exists := ctx.Get("user")
	if !exists {
		ctx.AbortWithStatus(http.StatusUnauthorized)
		return
	}
	input.UserName = user_name.(string)

	savedProduct, err := con.uc.SaveProduct(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, savedProduct)
}
