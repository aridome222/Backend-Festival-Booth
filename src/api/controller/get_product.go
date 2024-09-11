package controller

import "github.com/aridome222/Backend-Festival-Booth/usecase"

type GetProductController struct {
	uc usecase.GetProductUseCase
}

func NewGetProductController(uc usecase.GetProductUseCase) GetProductController {
	return GetProductController{
		uc: uc,
	}
}
