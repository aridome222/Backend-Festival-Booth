package controller

import (
	"net/http"

	"github.com/aridome222/Backend-Festival-Booth/usecase"
	"github.com/gin-gonic/gin"
)

type SaveAnswerController struct {
	uc usecase.SaveAnswerUseCase
}

func NewSaveAnswerController(uc usecase.SaveAnswerUseCase) SaveAnswerController {
	return SaveAnswerController{
		uc: uc,
	}
}

func (con SaveAnswerController) SaveAnswer(ctx *gin.Context) {
	var input usecase.SaveAnswerInputDTO

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

	savedAccount, err := con.uc.SaveAnswer(input)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, savedAccount)
}
