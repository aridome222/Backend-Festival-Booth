package usecase

import "github.com/aridome222/Backend-Festival-Booth/domain"

type GetProfileListUseCase struct {
	repo domain.ProfileRepository
}

type GetProfileListUseCaseInputDTO struct {
	Page  int
	Limit int
}

type GetProfileListUseCaseOutputDTO struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Introduction string `json:"introduction"`
	IconNum      int    `json:"icon_num"`
	GithubUrl    string `json:"github_url"`
	XUrl         string `json:"x_url"`
}

func NewGetProfileListUseCase(repo domain.ProfileRepository) GetProfileListUseCase {
	return GetProfileListUseCase{
		repo: repo,
	}
}

func (uc GetProfileListUseCase) GetProfileList(input GetProfileListUseCaseInputDTO) ([]GetProfileListUseCaseOutputDTO, error) {
	var profiles []domain.Profile
	var err error

	if input.Page < 0 || input.Limit < 0 {
		profiles, err = uc.repo.FindAll()
	} else {
		profiles, err = uc.repo.Find(input.Page, input.Limit)
	}

	if err != nil {
		return nil, err
	}

	outputSlice := []GetProfileListUseCaseOutputDTO{}

	for _, profile := range profiles {
		outputSlice = append(outputSlice, GetProfileListUseCaseOutputDTO{
			profile.ID,
			profile.Name,
			profile.Introduction,
			profile.IconNum,
			profile.GithubUrl,
			profile.XUrl,
		})
	}

	return outputSlice, nil
}
