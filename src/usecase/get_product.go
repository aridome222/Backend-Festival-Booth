package usecase

import "github.com/aridome222/Backend-Festival-Booth/domain"

type GetProductUseCase struct {
	repo domain.Product
}

func NewGetProductUseCase(repo domain.Product) GetProductUseCase {
	return GetProductUseCase{
		repo: repo,
	}
}
