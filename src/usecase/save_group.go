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

}
