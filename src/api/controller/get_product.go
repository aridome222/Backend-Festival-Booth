package controller

import (
	"net/http"

	"github.com/aridome222/Backend-Festival-Booth/usecase"
	"github.com/gin-gonic/gin"
)

type GetProductController struct {
	uc usecase.GetProductUseCase
}

func NewGetProductController(uc usecase.GetProductUseCase) GetProductController {
	return GetProductController{
		uc: uc,
	}
}

func (con GetProductController) GetProduct(ctx *gin.Context) {
	var input usecase.GetProductUseCaseInputDTO
	var err error

	input.UserName = ctx.Param("user_name")

	output, err := con.uc.GetProduct(input)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, output)
}
