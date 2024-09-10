package usecase

import "github.com/aridome222/Backend-Festival-Booth/domain"

type GetUserUseCase struct {
	repo domain.UserRepository
}

func NewGetUserUseCase(repo domain.UserRepository) GetUserUseCase {
	return GetUserUseCase{
		repo: repo,
	}
}
