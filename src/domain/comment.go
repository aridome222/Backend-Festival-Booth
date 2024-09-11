package domain

type Comment struct {
	ID        string `gorm:"primary_key"`
	ProductID string `gorm:"not null"`
	Message   string `gorm:"not null"`
}

type CommentRepository interface {
	Save(Comment) (Comment, error)
}

func NewComment(
	id string,
	productID string,
	message string,
) Comment {
	return Comment{
		ID:        id,
		ProductID: productID,
		Message:   message,
	}
}
