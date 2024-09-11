package usecase

import "github.com/aridome222/Backend-Festival-Booth/infrastructure/repository"

type GetProductUseCase struct {
	repo repository.ProductRepository
}

func NewGetProductUseCase(repo repository.ProductRepository) GetProductUseCase {
	return GetProductUseCase{
		repo: repo,
	}
}
