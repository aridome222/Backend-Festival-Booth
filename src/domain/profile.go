package domain

type Profile struct {
	ID           string `gorm:"primary_key"`
	Name         string `gorm:"unique;not null"`
	Introduction string `gorm:"not null"`
	IconNum      int    `gorm:"not null"`
	GithubUrl    string
	XUrl         string
}

type ProfileRepository interface {
	Save(Profile) (Profile, error)
	Find(int, int) ([]Profile, error)
	FindAll() ([]Profile, error)
	FindByUser(string) (Profile, error)
}

func NewProfile(
	id string,
	name string,
	introduction string,
	iconNum int,
	githubUrl string,
	xUrl string,
) Profile {
	return Profile{
		ID:           id,
		Name:         name,
		Introduction: introduction,
		IconNum:      iconNum,
		GithubUrl:    githubUrl,
		XUrl:         xUrl,
	}
}

// func (profile Profile) ID() string {
// 	return profile.id
// }

// func (profile Profile) Name() string {
// 	return profile.name
// }

// func (profile Profile) Introduction() string {
// 	return profile.introduction
// }

// func (profile Profile) IconNum() int {
// 	return profile.iconNum
// }

// func (profile Profile) GithubUrl() string {
// 	return profile.githubUrl
// }

// func (profile Profile) XUrl() string {
// 	return profile.xUrl
// }
