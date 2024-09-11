package controller

import (
	"net/http"

	"github.com/aridome222/Backend-Festival-Booth/usecase"
	"github.com/gin-gonic/gin"
)

type SaveCommentController struct {
	uc usecase.SaveCommentUseCase
}

func NewSaveCommentController(uc usecase.SaveCommentUseCase) SaveCommentController {
	return SaveCommentController{
		uc: uc,
	}
}

func (con SaveCommentController) SaveComment(ctx *gin.Context) {
	var input usecase.SaveCommentUseCaseInputDTO

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	savedComment, err := con.uc.SaveComment(input)
	// TODO: errがproductIDが一致するproductが存在しない場合は、400で返したい
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, savedComment)
}
