package controller

import (
	"net/http"

	"github.com/aridome222/Backend-Festival-Booth/usecase"
	"github.com/gin-gonic/gin"
)

type GetCommentController struct {
	uc usecase.GetCommentUseCase
}

func NewGetCommentController(uc usecase.GetCommentUseCase) GetCommentController {
	return GetCommentController{
		uc: uc,
	}
}

func (con GetCommentController) GetComment(ctx *gin.Context) {
	var input usecase.GetCommentInputDTO

	input.ProductID = ctx.Param("product_id")

	output, err := con.uc.GetComment(input)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, output)
}
