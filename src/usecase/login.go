package usecase

import "github.com/aridome222/Backend-Festival-Booth/domain"

type LoginUseCase struct {
	repo domain.AccountRepository
}

func NewLoginUseCase(repo domain.AccountRepository) LoginUseCase {
	return LoginUseCase{
		repo: repo,
	}
}
