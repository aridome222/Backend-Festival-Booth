package usecase

import "github.com/aridome222/Backend-Festival-Booth/domain"

type GetProfileUseCase struct {
	repo domain.ProfileRepository
}

type GetProfileUseCaseInputDTO struct {
	Page  int
	Limit int
}

type GetProfileUseCaseOutputDTO struct {
	ID           string
	Name         string
	Introduction string
	IconNum      int
	GithubUrl    string
	XUrl         string
}

func NewGetProfileUseCase(repo domain.ProfileRepository) GetProfileUseCase {
	return GetProfileUseCase{
		repo: repo,
	}
}

func (uc GetProfileUseCase) GetProfile(input GetProfileUseCaseInputDTO) ([]GetProfileUseCaseOutputDTO, error) {
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

	outputSlice := []GetProfileUseCaseOutputDTO{}

	for _, profile := range profiles {
		outputSlice = append(outputSlice, GetProfileUseCaseOutputDTO{
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
