package controller

import "github.com/aridome222/Backend-Festival-Booth/usecase"

type SaveProfileController struct {
	uc usecase.SaveProfileUsecase
}

func NewSaveProfileController(uc usecase.SaveProfileUsecase) SaveProfileController {
	return SaveProfileController{
		uc: uc,
	}
}
