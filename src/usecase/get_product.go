package usecase

import "github.com/aridome222/Backend-Festival-Booth/infrastructure/repository"

type GetProductUseCase struct {
	repo repository.ProductRepository
}

type GetProductUseCaseInputDTO struct {
	Page  int
	Limit int
}

type GetProductUseCaseOutputDTO struct {
	ProductID   string `json:"product_id"`
	UserName    string `json:"user_name"`
	Url         string `json:"url"`
	Description string `json:"description"`
}

func NewGetProductUseCase(repo repository.ProductRepository) GetProductUseCase {
	return GetProductUseCase{
		repo: repo,
	}
}
