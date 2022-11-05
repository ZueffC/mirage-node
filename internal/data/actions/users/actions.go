package users

import (
	"errors"

	"github.com/zueffc/mirage-node/internal/data"
	"github.com/zueffc/mirage-node/pkg/checker"
)

func Create(nick, email, password string) error {
	if checker.SizeCheck(nick, 3, 96) && checker.SizeCheck(email, 9, 128) && checker.SizeCheck(password, 128, 128) {

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
}

func Edit(id int) {}
