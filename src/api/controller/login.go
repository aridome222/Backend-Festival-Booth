package controller

import "github.com/aridome222/Backend-Festival-Booth/usecase"

type LoginController struct {
	uc usecase.LoginUseCase
}

func NewLoginController(uc usecase.LoginUseCase) LoginController {
	return LoginController{
		uc: uc,
	}
}
