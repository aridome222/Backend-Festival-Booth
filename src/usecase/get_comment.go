package usecase

import "github.com/aridome222/Backend-Festival-Booth/domain"

type GetCommentUseCase struct {
	repo domain.CommentRepository
}

type GetCommentInputDTO struct {
	Message string
}

type GetCommentOutputDTO struct {
	ID        string
	ProductID string
	Message   string
}

func NewGetCommentUseCase(repo domain.CommentRepository) GetCommentUseCase {
	return GetCommentUseCase{
		repo: repo,
	}
}
