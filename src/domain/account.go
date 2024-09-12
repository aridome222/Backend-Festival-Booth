package domain

type Account struct {
	ID       string `gorm:"primary_key"`
	Name     string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
}

type AccountRepository struct {
	// TODO
}

func NewAccount(
	id string,
	name string,
	password string,
) Account {
	return Account{
		ID:       id,
		Name:     name,
		Password: password,
	}
}
