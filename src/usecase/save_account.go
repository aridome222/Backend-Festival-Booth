package usecase

import "github.com/aridome222/Backend-Festival-Booth/domain"

type SaveAccountUseCase struct {
	repo domain.AccountRepository
}

type SaveAccountInputDTO struct {
	UserName string `json:"user_name" binding:"required,min=1,max=20"`
	Password string `json:"password" binding:"required,min=1"`
}

type SaveAccountOutputDTO struct {
	ID       string `json:"id"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

func NewSaveAccountUseCase(repo domain.AccountRepository) SaveAccountUseCase {
	return SaveAccountUseCase{
		repo: repo,
	}
}
