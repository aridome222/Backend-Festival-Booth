package usecase

import "github.com/aridome222/Backend-Festival-Booth/domain"

type LoginUseCase struct {
	repo domain.AccountRepository
}

type LoginInputDTO struct {
	UserName string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func NewLoginUseCase(repo domain.AccountRepository) LoginUseCase {
	return LoginUseCase{
		repo: repo,
	}
}

/*
 * ここではトークンの生成は行わず、ログインの成否だけを判定する
 * これにより、認証手段を変更しやすくしている
 */
func (uc LoginUseCase) Login(input LoginInputDTO) (string, bool) {
	account, err := uc.repo.FindByName(input.UserName)
	if err != nil {
		return "", false
	}

	return account.Name, account.IsValidPassword(input.Password)
}
