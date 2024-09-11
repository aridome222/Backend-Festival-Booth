package controller

import "github.com/aridome222/Backend-Festival-Booth/usecase"

type GetProfileController struct {
	uc usecase.GetProfileUseCase
}

func NewGetProfileController(uc usecase.GetProfileUseCase) GetProfileController {
	return GetProfileController{
		uc: uc,
	}
}
