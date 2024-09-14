package usecase

import "github.com/aridome222/Backend-Festival-Booth/domain"

type SaveAnswerUseCase struct {
	repo domain.AccountRepository
}

type SaveAnswerInputDTO struct {
	UserName string
	Answer   int `json:"answer"`
}

type SaveAnswerOutputDTO struct {
	UserName string `json:"name"`
	Answer   int    `json:"answer"`
}

func NewSaveAnswerUseCase(repo domain.AccountRepository) SaveAnswerUseCase {
	return SaveAnswerUseCase{
		repo: repo,
	}
}

func (uc SaveAnswerUseCase) SaveAnswer(input SaveAnswerInputDTO) (SaveAnswerOutputDTO, error) {
	account, err := uc.repo.FindByName(input.UserName)
	// TODO: accountが見つからなかった場合以外のエラーハンドリングを記述
	if err != nil {
		return SaveAnswerOutputDTO{}, err
	}

	account, err = domain.NewAccount(account.ID, account.Name, account.Password, input.Answer)
	if err != nil {
		return SaveAnswerOutputDTO{}, err
	}

	account, err = uc.repo.Save(account)
	if err != nil {
		return SaveAnswerOutputDTO{}, err
	}

	return SaveAnswerOutputDTO{
		account.Name,
		account.Answer,
	}, nil
}
