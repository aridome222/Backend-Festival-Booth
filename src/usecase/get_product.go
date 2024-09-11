package usecase

import "github.com/aridome222/Backend-Festival-Booth/domain"

type GetProductUseCase struct {
	repo domain.Product
}

type GetProductUseCaseInputDTO struct {
	UserName string
}

type GetProductUseCaseOutputDTO struct {
	ProductID   string `json:"product_id"`
	UserName    string `json:"user_name"`
	Url         string `json:"url"`
	Description string `json:"description"`
}

func NewGetProductUseCase(repo domain.Product) GetProductUseCase {
	return GetProductUseCase{
		repo: repo,
	}
}
