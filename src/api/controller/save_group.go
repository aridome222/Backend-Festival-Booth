package controller

import "github.com/aridome222/Backend-Festival-Booth/usecase"

type SaveGroupController struct {
	uc usecase.SaveGroupUseCase
}

func NewGroupController(uc usecase.SaveGroupUseCase) SaveGroupController {
	return SaveGroupController{
		uc: uc,
	}
}
