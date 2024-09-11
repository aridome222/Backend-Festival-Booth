package usecase

import "github.com/aridome222/Backend-Festival-Booth/domain"

type SaveCommentUseCase struct {
	repo domain.CommentRepository
}

type SaveCommentUseCaseInputDTO struct {
	ProductID string `json:"product_id" binding:"required"`
	Message   string `json:"message" binding:"required,min=1,max=100"`
}

type SaveCommentUseCaseOutputDTO struct {
	ID        string `json:"id"`
	ProductID string `json:"product_id"`
	Message   string `json:"message"`
}

func NewCommentUseCase(repo domain.CommentRepository) SaveCommentUseCase {
	return SaveCommentUseCase{
		repo: repo,
	}
}
