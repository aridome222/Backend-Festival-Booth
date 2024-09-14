package domain

import "golang.org/x/crypto/bcrypt"

type Account struct {
	ID       string `gorm:"primary_key"`
	Name     string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Group    int
}

type AccountRepository interface {
	Create(Account) (Account, error)
	Save(Account) (Account, error)
	FindByName(string) (Account, error)
}

func NewAccount(
	id string,
	name string,
	password string,
	group int,
) (Account, error) {
	// パスワードをハッシュ化
	hash_pass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return Account{}, err
	}

	return Account{
		ID:       id,
		Name:     name,
		Password: string(hash_pass),
		Group:    group,
	}, nil
}

func (account Account) IsValidPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(account.Password), []byte(password))
	return err != bcrypt.ErrMismatchedHashAndPassword
}
