package controller

import (
	"net/http"
	"strconv"

	"github.com/aridome222/Backend-Festival-Booth/usecase"
	"github.com/gin-gonic/gin"
)

type GetProfileController struct {
	uc usecase.GetProfileUseCase
}

func NewGetProfileController(uc usecase.GetProfileUseCase) GetProfileController {
	return GetProfileController{
		uc: uc,
	}
}

func (con GetProfileController) GetProfile(ctx *gin.Context) {
	var input usecase.GetProfileUseCaseInputDTO
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

	output, err := con.uc.GetProfile(input)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, output)
}
