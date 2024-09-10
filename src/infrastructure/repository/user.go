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

func (repo UserRepository) Find(page int, limit int) ([]domain.User, error) {
	var users []domain.User
	result := repo.db.Limit(limit).Offset(page * limit).Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}

func (repo UserRepository) FindAll() ([]domain.User, error) {
	var users []domain.User
	result := repo.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}
