package usecase

import (
	"github.com/aridome222/Backend-Festival-Booth/domain"
)

type GetProductListUseCase struct {
	repo domain.ProductRepository
}

type GetProductListUseCaseInputDTO struct {
	Page  int
	Limit int
}

type GetProductListUseCaseOutputDTO struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	UserName    string `json:"user_name"`
	Url         string `json:"url"`
	Description string `json:"description"`
}

func NewGetProductListUseCase(repo domain.ProductRepository) GetProductListUseCase {
	return GetProductListUseCase{
		repo: repo,
	}
}

func (uc GetProductListUseCase) GetProductList(input GetProductListUseCaseInputDTO) ([]GetProductListUseCaseOutputDTO, error) {
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

	outputSlice := []GetProductListUseCaseOutputDTO{}

	for _, product := range products {
		outputSlice = append(outputSlice, GetProductListUseCaseOutputDTO{
			product.ID,
			product.Title,
			product.UserName,
			product.Url,
			product.Description,
		})
	}

	return outputSlice, nil
}
