package usecase

import "github.com/aridome222/Backend-Festival-Booth/domain"

type SaveGroupUseCase struct {
	repo domain.AccountRepository
}

type SaveGroupInputDTO struct {
	UserName string
	Group    int `json:"group"`
}

type SaveGroupOutputDTO struct {
	UserName string `json:"name"`
	Group    int    `json:"group"`
}

func NewSaveGroupUseCase(repo domain.AccountRepository) SaveGroupUseCase {
	return SaveGroupUseCase{
		repo: repo,
	}
}

func (uc SaveGroupUseCase) SaveGroup(input SaveGroupInputDTO) (SaveGroupOutputDTO, error) {
	account, err := uc.repo.FindByName(input.UserName)
	// TODO: accountが見つからなかった場合以外のエラーハンドリングを記述
	if err != nil {
		return SaveGroupOutputDTO{}, err
	}

	account, err = domain.NewAccount(account.ID, account.Name, account.Password, input.Group)
	if err != nil {
		return SaveGroupOutputDTO{}, err
	}

	account, err = uc.repo.Save(account)
	if err != nil {
		return SaveGroupOutputDTO{}, err
	}

	return SaveGroupOutputDTO{
		account.Name,
		account.Group,
	}, nil
}
