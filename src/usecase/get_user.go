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
