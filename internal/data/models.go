package data

import (
	pkgs "github.com/zueffc/mirage-node/internal/data/actions/packages"
	"gorm.io/gorm"
)

type User struct {
	Nick     string `json:"nick" gorm:"unique"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
}

type UserModel struct {
	gorm.Model
	User
}

type PackageModel struct {
	gorm.Model
	pkgs.Package
}
