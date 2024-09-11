package usecase

import "github.com/aridome222/Backend-Festival-Booth/domain"

type SaveCommentUseCase struct {
	repo domain.CommentRepository
}

func NewCommentUseCase(repo domain.CommentRepository) SaveCommentUseCase {
	return SaveCommentUseCase{
		repo: repo,
	}
}
