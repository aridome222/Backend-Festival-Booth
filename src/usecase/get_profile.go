package usecase

import "github.com/aridome222/Backend-Festival-Booth/domain"

type GetProfileUseCase struct {
	repo domain.ProfileRepository
}

func NewGetProfileUseCase(repo domain.ProfileRepository) GetProfileUseCase {
	return GetProfileUseCase{
		repo: repo,
	}
}
