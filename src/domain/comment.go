package domain

type Comment struct {
	ID        string `gorm:"primary_key"`
	ProductID string `gorm:"not null"`
	UserName  string `gorm:"not null"`
	Message   string `gorm:"not null"`
}

type CommentRepository interface {
}

func NewComment(
	id string,
	productID string,
	userName string,
	message string,
) Comment {
	return Comment{
		ID:        id,
		ProductID: productID,
		UserName:  userName,
		Message:   message,
	}
}
