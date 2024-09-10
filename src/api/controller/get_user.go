package controller

import "github.com/aridome222/Backend-Festival-Booth/usecase"

type GetUserController struct {
	uc usecase.GetUserUseCase
}

func NewGetUserController(uc usecase.GetUserUseCase) GetUserController {
	return GetUserController{
		uc: uc,
	}
}
