package domain

type User struct {
	ID           string `gorm:"primary_key"`
	Name         string `gorm:"unique;not null"`
	Introduction string `gorm:"not null"`
	IconNum      int    `gorm:"not null"`
	GithubUrl    string
	XUrl         string
}

type UserRepository interface {
	// Save(User) (User, error)
	Find(int, int) ([]User, error)
	FindAll() ([]User, error)
}

func NewUser(
	id string,
	name string,
	introduction string,
	iconNum int,
	githubUrl string,
	xUrl string,
) User {
	return User{
		ID:           id,
		Name:         name,
		Introduction: introduction,
		IconNum:      iconNum,
		GithubUrl:    githubUrl,
		XUrl:         xUrl,
	}
}

// func (user User) ID() string {
// 	return user.id
// }

// func (user User) Name() string {
// 	return user.name
// }

// func (user User) Introduction() string {
// 	return user.introduction
// }

// func (user User) IconNum() int {
// 	return user.iconNum
// }

// func (user User) GithubUrl() string {
// 	return user.githubUrl
// }

// func (user User) XUrl() string {
// 	return user.xUrl
// }
