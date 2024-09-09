package domain

type Product struct {
	productID   string
	userName    string
	url         string
	description string
}

type ProductRepository interface {
	Save(Product) (Product, error)
	FindByUser(string) (Product, error)
}

func NewProduct(id string, userName string, url string, description string) Product {
	return Product{
		productID:   id,
		userName:    userName,
		url:         url,
		description: description,
	}
}

func (p Product) ProductID() string {
	return p.productID
}

func (p Product) UserName() string {
	return p.userName
}

func (p Product) Url() string {
	return p.url
}

func (p Product) Description() string {
	return p.description
}
