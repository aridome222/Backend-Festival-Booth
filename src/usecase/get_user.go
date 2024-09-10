package usecase

import "github.com/aridome222/Backend-Festival-Booth/domain"

type GetUserUseCase struct {
	repo domain.UserRepository
}

type GetUserUseCaseInputDTO struct {
	Page  int
	Limit int
}

type GetUserUseCaseOutputDTO struct {
	ID           string
	Name         string
	Introduction string
	IconNum      int
	GithubUrl    string
	XUrl         string
}

func NewGetUserUseCase(repo domain.UserRepository) GetUserUseCase {
	return GetUserUseCase{
		repo: repo,
	}
}

func (uc GetUserUseCase) GetUser(input GetUserUseCaseInputDTO) ([]GetUserUseCaseOutputDTO, error) {
	var users []domain.User
	var err error

	if input.Page < 0 || input.Limit <= 0 {
		users, err = uc.repo.FindAll()
	} else {
		users, err = uc.repo.Find(input.Page, input.Limit)
	}

	if err != nil {
		return nil, err
	}

	outputSlice := []GetUserUseCaseOutputDTO{}

	for _, user := range users {
		outputSlice = append(outputSlice, GetUserUseCaseOutputDTO{
			user.ID,
			user.Name,
			user.Introduction,
			user.IconNum,
			user.GithubUrl,
			user.XUrl,
		})
	}

	return outputSlice, nil
}
