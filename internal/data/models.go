package data

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	Nick     string `json:"nick" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

type Package struct {
	AuthorID    uint   `json:"author_id"`
	Name        string `json:"name" gorm:"unique"`
	Description string `json:"description" gorm:"unique"`
	GitURL      string `json:"git_url" gorm:"unique"`
}

type UserModel struct {
	gorm.Model
	User
}

type PackageModel struct {
	gorm.Model
	Package
}
