package controller

import "github.com/aridome222/Backend-Festival-Booth/usecase"

type SaveCommentController struct {
	uc usecase.SaveCommentUseCase
}

func NewCommentController(uc usecase.SaveCommentUseCase) SaveCommentController {
	return SaveCommentController{
		uc: uc,
	}
}
