package usecase

import "github.com/aridome222/Backend-Festival-Booth/domain"

type GetProductUseCase struct {
	repo domain.ProductRepository
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

func NewGetProductUseCase(repo domain.ProductRepository) GetProductUseCase {
	return GetProductUseCase{
		repo: repo,
	}
}

func (uc GetProductUseCase) GetProduct(input GetProductUseCaseInputDTO) (GetProductUseCaseOutputDTO, error) {
	var product domain.Product
	var err error

	product, err = uc.repo.FindByUser(input.UserName)
	if err != nil {
		return GetProductUseCaseOutputDTO{}, err
	}

	return GetProductUseCaseOutputDTO{
		product.ProductID,
		product.UserName,
		product.Url,
		product.Description,
	}, nil
}
