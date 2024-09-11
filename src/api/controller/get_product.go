package controller

import (
	"net/http"
	"strconv"

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

	input.Page, err = strconv.Atoi(ctx.DefaultQuery("page", "-1"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	input.Limit, err = strconv.Atoi(ctx.DefaultQuery("limit", "-1"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	output, err := con.uc.GetProduct(input)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, output)
}
