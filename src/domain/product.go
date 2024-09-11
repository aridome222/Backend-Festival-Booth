package domain

type Product struct {
	ID          string `gorm:"primary_key"`
	Title       string `gorm:"not null"`
	UserName    string `gorm:"unique;not null"`
	Url         string `gorm:"not null"`
	Description string
}

type ProductRepository interface {
	Save(Product) (Product, error)
	Find(int, int) ([]Product, error)
	FindAll() ([]Product, error)
	FindByID(string) (Product, error)
	FindByUser(string) (Product, error)
}

func NewProduct(
	id string,
	title string,
	userName string,
	url string,
	description string,
) Product {
	return Product{
		ID:          id,
		Title:       title,
		UserName:    userName,
		Url:         url,
		Description: description,
	}
}

// func (p Product) ProductID() string {
// 	return p.productID
// }

// func (p Product) UserName() string {
// 	return p.userName
// }

// func (p Product) Url() string {
// 	return p.url
// }

// func (p Product) Description() string {
// 	return p.description
// }
