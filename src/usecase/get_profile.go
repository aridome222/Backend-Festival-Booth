package usecase

import "github.com/aridome222/Backend-Festival-Booth/domain"

type GetProfileUseCase struct {
	repo domain.ProfileRepository
}

type GetProfileUseCaseInputDTO struct {
	UserName string
}

type GetProfileUseCaseOutputDTO struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Introduction string `json:"introduction"`
	IconNum      int    `json:"icon_num"`
	GithubUrl    string `json:"github_url"`
	XUrl         string `json:"x_url"`
}

func NewGetProfileUseCase(repo domain.ProfileRepository) GetProfileUseCase {
	return GetProfileUseCase{
		repo: repo,
	}
}
