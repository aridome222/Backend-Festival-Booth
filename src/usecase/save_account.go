package usecase

import "github.com/aridome222/Backend-Festival-Booth/domain"

type SaveAccountUseCase struct {
	repo domain.AccountRepository
}

func NewSaveAccountUseCase(repo domain.AccountRepository) SaveAccountUseCase {
	return SaveAccountUseCase{
		repo: repo,
	}
}
