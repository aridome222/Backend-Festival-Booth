package usecase

import (
	"github.com/aridome222/Backend-Festival-Booth/domain"
	"github.com/oklog/ulid/v2"
)

type SaveProfileUsecase struct {
	repo        domain.ProfileRepository
	accountRepo domain.AccountRepository
}

type SaveProfileInputDTO struct {
	// Name         string `json:"name"`
	// Introduction string `json:"introduction"`
	// IconNum      int    `json:"icon_num"`
	// GithubUrl    string `json:"github_url"`
	// XUrl         string `json:"x_url"`
	Name         string
	Introduction string `json:"introduction" binding:"required,min=1,max=200"`
	IconNum      int    `json:"icon_num"`
	GithubUrl    string `json:"github_url"`
	XUrl         string `json:"x_url"`
}

type SaveProfileOutputDTO struct {
	ID           string `json:"id"`
	Name         string `json:"name"`
	Introduction string `json:"introduction"`
	IconNum      int    `json:"icon_num"`
	GithubUrl    string `json:"github_url"`
	XUrl         string `json:"x_url"`
}

func NewSaveProfileUseCase(
	repo domain.ProfileRepository,
	accountRepo domain.AccountRepository,
) SaveProfileUsecase {
	return SaveProfileUsecase{
		repo:        repo,
		accountRepo: accountRepo,
	}
}

func (uc SaveProfileUsecase) SaveProfile(input SaveProfileInputDTO) (SaveProfileOutputDTO, error) {
	_, err := uc.accountRepo.FindByName(input.Name)
	// TODO: accountが見つからなかった場合以外のエラーハンドリングを記述
	if err != nil {
		return SaveProfileOutputDTO{}, err
	}

	profile, err := uc.repo.FindByUser(input.Name)
	// TODO: レコードが何も見つからなった場合以外のエラーハンドリングを記述する
	// if err != nil {
	// 	return SaveProfileOutputDTO{}, err
	// }

	if err != nil {
		// 新規追加
		profile = domain.NewProfile(ulid.Make().String(), input.Name, input.Introduction, input.IconNum, input.GithubUrl, input.XUrl)
	} else {
		//　更新
		profile = domain.NewProfile(profile.ID, input.Name, input.Introduction, input.IconNum, input.GithubUrl, input.XUrl)
	}

	profile, err = uc.repo.Save(profile)

	if err != nil {
		return SaveProfileOutputDTO{}, err
	}

	return SaveProfileOutputDTO{
		ID:           profile.ID,
		Name:         profile.Name,
		Introduction: profile.Introduction,
		IconNum:      profile.IconNum,
		GithubUrl:    profile.GithubUrl,
		XUrl:         profile.XUrl,
	}, nil
}
