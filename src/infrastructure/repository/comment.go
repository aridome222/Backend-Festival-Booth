package repository

import (
	"github.com/aridome222/Backend-Festival-Booth/domain"
	"gorm.io/gorm"
)

type CommentRepository struct {
	db *gorm.DB
}

func NewCommentRepository(db *gorm.DB) domain.CommentRepository {
	return CommentRepository{
		db: db,
	}
}

func (repo CommentRepository) Save(comment domain.Comment) (domain.Comment, error) {
	result := repo.db.Save(comment)
	if result.Error != nil {
		return domain.Comment{}, result.Error
	}

	return comment, nil
}
