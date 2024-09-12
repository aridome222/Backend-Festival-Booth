package usecase

import "github.com/aridome222/Backend-Festival-Booth/domain"

type LoginUseCase struct {
	repo domain.AccountRepository
}

type LoginInputDTO struct {
	UserName string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginOutputDTO struct {
	UserName string `json:"name"`
}

func NewLoginUseCase(repo domain.AccountRepository) LoginUseCase {
	return LoginUseCase{
		repo: repo,
	}
}
