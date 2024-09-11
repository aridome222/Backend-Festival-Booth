package repository

import (
	"github.com/aridome222/Backend-Festival-Booth/domain"
	"gorm.io/gorm"
)

type ProfileRepository struct {
	db *gorm.DB
}

func NewProfileRepository(db *gorm.DB) domain.ProfileRepository {
	return ProfileRepository{
		db: db,
	}
}

func (repo ProfileRepository) Save(profile domain.Profile) (domain.Profile, error) {
	result := repo.db.Save(profile)
	if result.Error != nil {
		return domain.Profile{}, result.Error
	}
	return profile, nil
}

func (repo ProfileRepository) Find(page int, limit int) ([]domain.Profile, error) {
	var profiles []domain.Profile
	result := repo.db.Limit(limit).Offset(page * limit).Find(&profiles)
	if result.Error != nil {
		return nil, result.Error
	}
	return profiles, nil
}

func (repo ProfileRepository) FindAll() ([]domain.Profile, error) {
	var profiles []domain.Profile
	result := repo.db.Find(&profiles)
	if result.Error != nil {
		return nil, result.Error
	}
	return profiles, nil
}

func (repo ProfileRepository) FindByUser(name string) (domain.Profile, error) {
	var profile domain.Profile

	result := repo.db.Where("name = ?", name).First(&profile)
	if result.Error != nil {
		return domain.Profile{}, result.Error
	}

	return profile, nil
}
