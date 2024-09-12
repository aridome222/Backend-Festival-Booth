package usecase

import "github.com/aridome222/Backend-Festival-Booth/domain"

type GetCommentUseCase struct {
	repo domain.CommentRepository
}

type GetCommentInputDTO struct {
	ProductID string `json:"product_id" binding:"required"`
	Message   string `json:"message" binding:"required,min=1,max=100"`
}

type GetCommentOutputDTO struct {
	ID        string `json:"id"`
	ProductID string `json:"product_id"`
	Message   string `json:"message"`
}

func NewGetCommentUseCase(repo domain.CommentRepository) GetCommentUseCase {
	return GetCommentUseCase{
		repo: repo,
	}
}
