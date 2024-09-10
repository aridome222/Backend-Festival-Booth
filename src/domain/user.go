package domain

type User struct {
	id           string
	name         string
	introduction string
	iconNum      string
	githubUrl    string
	xUrl         string
}

func NewUser(
	id string,
	name string,
	introduction string,
	iconNum string,
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
