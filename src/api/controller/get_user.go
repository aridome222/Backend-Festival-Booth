package controller

import (
	"net/http"
	"strconv"

	"github.com/aridome222/Backend-Festival-Booth/usecase"
	"github.com/gin-gonic/gin"
)

type GetUserController struct {
	uc usecase.GetUserUseCase
}

func NewGetUserController(uc usecase.GetUserUseCase) GetUserController {
	return GetUserController{
		uc: uc,
	}
}

func (con GetUserController) GetUser(ctx *gin.Context) {
	var input usecase.GetUserUseCaseInputDTO

	input.Page, _ = strconv.Atoi(ctx.DefaultQuery("page", "-1"))
	input.Limit, _ = strconv.Atoi(ctx.DefaultQuery("limit", "-1"))

	output, err := con.uc.GetUser(input)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, output)
}
