package controller

import "github.com/aridome222/Backend-Festival-Booth/usecase"

type GetCommentController struct {
	uc usecase.GetCommentUseCase
}

func NewGetCommentController(uc usecase.GetCommentUseCase) GetCommentController {
	return GetCommentController{
		uc: uc,
	}
}
