package repository

import (
	"github.com/aridome222/Backend-Festival-Booth/domain"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) domain.UserRepository {
	return UserRepository{
		db: db,
	}
}
