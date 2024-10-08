package domain

type Comment struct {
	ID        string `gorm:"primary_key"`
	UserName  string `gorm:"not null"`
	ProductID string `gorm:"not null"`
	Message   string `gorm:"not null"`
}

type CommentRepository interface {
	Save(Comment) (Comment, error)
	FindByProductID(string) ([]Comment, error)
}

func NewComment(
	id string,
	user_name string,
	productID string,
	message string,
) Comment {
	return Comment{
		ID:        id,
		UserName:  user_name,
		ProductID: productID,
		Message:   message,
	}
}
