package domain

type Product struct {
	ProductID   string `gorm:"primary_key"`
	UserName    string `gorm:"not null"`
	Url         string `gorm:"not null"`
	Description string
}

type ProductRepository interface {
	Save(Product) (Product, error)
	FindByUser(string) (Product, error)
}

func NewProduct(id string, userName string, url string, description string) Product {
	return Product{
		ProductID:   id,
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
