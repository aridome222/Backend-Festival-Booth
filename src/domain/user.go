package domain

type User struct {
	id           string
	name         string
	introduction string
	iconNum      int
	githubUrl    string
	xUrl         string
}

type UserRepository interface {
	Save(User) (User, error)
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
		id:           id,
		name:         name,
		introduction: introduction,
		iconNum:      iconNum,
		githubUrl:    githubUrl,
		xUrl:         xUrl,
	}
}

func (user User) ID() string {
	return user.id
}

func (user User) Name() string {
	return user.name
}

func (user User) Introduction() string {
	return user.introduction
}

func (user User) IconNum() int {
	return user.iconNum
}

func (user User) GithubUrl() string {
	return user.githubUrl
}

func (user User) XUrl() string {
	return user.xUrl
}
