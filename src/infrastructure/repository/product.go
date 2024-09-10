package repository

import (
	"github.com/aridome222/Backend-Festival-Booth/domain"
	"gorm.io/gorm"
)

type ProductRepository struct {
	db *gorm.DB
}

func NewProductRepository(db *gorm.DB) domain.ProductRepository {
	return ProductRepository{
		db: db,
	}
}

func (repo ProductRepository) Save(product domain.Product) (domain.Product, error) {
	result := repo.db.Save(product)
	if result.Error != nil {
		return domain.Product{}, result.Error
	}

	return product, nil
}

func (repo ProductRepository) FindByUser(userName string) (domain.Product, error) {
	return domain.Product{}, nil
}
