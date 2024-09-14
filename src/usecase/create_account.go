package usecase

import (
	"errors"

	"github.com/aridome222/Backend-Festival-Booth/domain"
	"github.com/oklog/ulid/v2"
)

type CreateAccountUseCase struct {
	repo domain.AccountRepository
}

type CreateAccountInputDTO struct {
	UserName string `json:"name" binding:"required,min=1,max=20"`
	Password string `json:"password" binding:"required,min=1"`
}

type CreateAccountOutputDTO struct {
	ID       string `json:"id"`
	UserName string `json:"name"`
	Password string `json:"password"`
}

func NewCreateAccountUseCase(repo domain.AccountRepository) CreateAccountUseCase {
	return CreateAccountUseCase{
		repo: repo,
	}
}

func (uc CreateAccountUseCase) CreateAccount(input CreateAccountInputDTO) (CreateAccountOutputDTO, error) {
	// 同一アカウントが存在しないか確認
	account, err := uc.repo.FindByName(input.UserName)

	// TODO: アカウントが存在しなかった場合以外のエラー処理を記述
	if err == nil {
		return CreateAccountOutputDTO{}, errors.New("this account name is used")
	}

	account, err = domain.NewAccount(ulid.Make().String(), input.UserName, input.Password, -1)
	if err != nil {
		return CreateAccountOutputDTO{}, err
	}

	// 同一アカウントが存在しない場合はDBに保存
	account, err = uc.repo.Create(account)

	if err != nil {
		return CreateAccountOutputDTO{}, err
	}

	return CreateAccountOutputDTO{
		account.ID,
		account.Name,
		account.Password,
	}, nil
}
