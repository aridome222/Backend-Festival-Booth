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

func (uc GetProfileUseCase) GetProfile(input GetProfileUseCaseInputDTO) (GetProfileUseCaseOutputDTO, error) {
	var profile domain.Profile
	var err error

	profile, err = uc.repo.FindByUser(input.UserName)
	if err != nil {
		return GetProfileUseCaseOutputDTO{}, err
	}

	return GetProfileUseCaseOutputDTO{
		profile.ID,
		profile.Name,
		profile.Introduction,
		profile.IconNum,
		profile.GithubUrl,
		profile.XUrl,
	}, nil
}
