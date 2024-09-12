package controller

import "github.com/aridome222/Backend-Festival-Booth/usecase"

type CreateAccountController struct {
	uc usecase.CreateAccountUseCase
}

func NewCreateAccountController(uc usecase.CreateAccountUseCase) CreateAccountController {
	return CreateAccountController{
		uc: uc,
	}
}
