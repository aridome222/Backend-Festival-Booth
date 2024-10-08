package usecase

import (
	"github.com/aridome222/Backend-Festival-Booth/domain"
	"github.com/oklog/ulid/v2"
)

type SaveProductUseCase struct {
	repo        domain.ProductRepository
	accountRepo domain.AccountRepository
}

type SaveProductInputDTO struct {
	Title       string `json:"title" binding:"required,min=1,max=50"`
	UserName    string
	Url         string `json:"url" binding:"required"`
	Description string `json:"description" binding:"min=0,max=200"`
}

type SaveProductOutputDTO struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	UserName    string `json:"user_name"`
	Url         string `json:"url"`
	Description string `json:"description"`
}

func NewSaveProductUseCase(
	repo domain.ProductRepository,
	accountRepo domain.AccountRepository,
) SaveProductUseCase {
	return SaveProductUseCase{
		repo:        repo,
		accountRepo: accountRepo,
	}
}

func (uc SaveProductUseCase) SaveProduct(input SaveProductInputDTO) (SaveProductOutputDTO, error) {
	_, err := uc.accountRepo.FindByName(input.UserName)
	// TODO: accountが見つからなかった場合以外のエラーハンドリングを記述
	if err != nil {
		return SaveProductOutputDTO{}, err
	}

	product, err := uc.repo.FindByUser(input.UserName)
	// TODO: レコードが何も見つからなった場合以外のエラーハンドリングを記述する
	// if err != nil {
	// 	return SaveProductOutputDTO{}, err
	// }

	if err != nil {
		product = domain.NewProduct(ulid.Make().String(), input.Title, input.UserName, input.Url, input.Description)
	} else {
		product = domain.NewProduct(product.ID, input.Title, input.UserName, input.Url, input.Description)
	}

	product, err = uc.repo.Save(product)
	if err != nil {
		return SaveProductOutputDTO{}, err
	}

	return SaveProductOutputDTO{
		ID:          product.ID,
		Title:       product.Title,
		UserName:    product.UserName,
		Url:         product.Url,
		Description: product.Description,
	}, nil
}
