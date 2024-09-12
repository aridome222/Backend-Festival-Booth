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
