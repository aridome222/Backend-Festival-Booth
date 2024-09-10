package usecase

import (
	"github.com/aridome222/Backend-Festival-Booth/domain"
	"github.com/oklog/ulid/v2"
)

type SaveProductUseCase struct {
	repo domain.ProductRepository
}

type SaveProductInputDTO struct {
	UserName    string `json:"user_name" bindng:"required:min=1:max=20"`
	Url         string `json:"url" bindng:"required:min=1:max=100"`
	Description string `json:"description" bindng:"min=1:max=200"`
}

type SaveProductOutputDTO struct {
	ProductID   string `json:"product_id"`
	UserName    string `json:"user_name"`
	Url         string `json:"url"`
	Description string `json:"description"`
}

func NewSaveProductUseCase(repo domain.ProductRepository) SaveProductUseCase {
	return SaveProductUseCase{
		repo: repo,
	}
}

func (uc SaveProductUseCase) SaveProduct(input SaveProductInputDTO) (SaveProductOutputDTO, error) {
	product, err := uc.repo.FindByUser(input.UserName)
	// TODO: レコードが何も見つからなった場合以外のエラーハンドリングを記述する
	// if err != nil {
	// 	return SaveProductOutputDTO{}, err
	// }

	if err != nil {
		product = domain.NewProduct(ulid.Make().String(), input.UserName, input.Url, input.Description)
	} else {
		product = domain.NewProduct(product.ProductID(), input.UserName, input.Url, input.Description)
	}

	product, err = uc.repo.Save(product)
	if err != nil {
		return SaveProductOutputDTO{}, err
	}

	return SaveProductOutputDTO{
		ProductID:   product.ProductID(),
		UserName:    product.UserName(),
		Url:         product.Url(),
		Description: product.Description(),
	}, nil
}
