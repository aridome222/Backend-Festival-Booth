package controller

import (
	"net/http"

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

	input.UserName = ctx.Param("name")

	output, err := con.uc.GetProfile(input)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, output)
}
