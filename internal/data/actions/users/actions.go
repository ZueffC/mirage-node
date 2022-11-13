package users

import (
	"errors"

	"github.com/zueffc/mirage-node/internal/data"
	"github.com/zueffc/mirage-node/pkg/checker"
)

func Create(nick, email, password string) error {
	if checker.SizeCheck(nick, 3, 96) && checker.SizeCheck(email, 9, 128) && checker.SizeCheck(password, 96, 96) {

		newUser := data.UserModel{
			User: data.User{
				Nick:     nick,
				Email:    email,
				Password: password,
			},
		}

		data.Database().Create(&newUser)
		return nil
	}

	return errors.New("Size is incompatible!")
}

func Delete(id int, password string) error {
	data.Database().Delete(&data.UserModel{}, id)

	return nil
}

func Edit(id int) {}

func Find(nick, password, email string) data.UserModel {
	var user data.UserModel

	error := data.Database().Where(&data.UserModel{
		User: data.User{
			Nick:     nick,
			Email:    email,
			Password: password,
		},
	}).Find(&user).Error

	if error != nil {
		return data.UserModel{}
	}
	return user
}
