package usecase

import "github.com/aridome222/Backend-Festival-Booth/domain"

type GetCommentUseCase struct {
	repo domain.CommentRepository
}

func NewGetCommentUseCase(repo domain.CommentRepository) GetCommentUseCase {
	return GetCommentUseCase{
		repo: repo,
	}
}
