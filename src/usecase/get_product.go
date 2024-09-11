package usecase

import (
	"github.com/aridome222/Backend-Festival-Booth/domain"
	"github.com/aridome222/Backend-Festival-Booth/infrastructure/repository"
)

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

func (uc GetProductUseCase) GetProduct(input GetProductUseCaseInputDTO) ([]GetProductUseCaseOutputDTO, error) {
	var products []domain.Product
	var err error

	if input.Page < 0 || input.Limit < 0 {
		products, err = uc.repo.FindAll()
	} else {
		products, err = uc.repo.Find(input.Page, input.Limit)
	}

	if err != nil {
		return nil, err
	}

	outputSlice := []GetProductUseCaseOutputDTO{}

	for _, product := range products {
		outputSlice = append(outputSlice, GetProductUseCaseOutputDTO{
			product.ProductID,
			product.UserName,
			product.Url,
			product.Description,
		})
	}

	return outputSlice, nil
}
