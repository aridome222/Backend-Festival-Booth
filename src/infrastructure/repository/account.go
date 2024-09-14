package repository

import (
	"github.com/aridome222/Backend-Festival-Booth/domain"
	"gorm.io/gorm"
)

type AccountRepository struct {
	db *gorm.DB
}

func NewAccountRepository(db *gorm.DB) domain.AccountRepository {
	return AccountRepository{
		db: db,
	}
}

func (repo AccountRepository) Create(account domain.Account) (domain.Account, error) {
	result := repo.db.Create(&account)
	if result.Error != nil {
		return domain.Account{}, result.Error
	}

	return account, nil
}

func (repo AccountRepository) Save(account domain.Account) (domain.Account, error) {
	result := repo.db.Save(account)
	if result.Error != nil {
		return domain.Account{}, result.Error
	}
	return account, nil
}

func (repo AccountRepository) FindByName(name string) (domain.Account, error) {
	var account domain.Account

	result := repo.db.Where("name = ?", name).First(&account)
	if result.Error != nil {
		return domain.Account{}, result.Error
	}

	return account, nil
}
