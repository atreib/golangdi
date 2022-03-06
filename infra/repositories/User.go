package repositories

import (
	"errors"

	"github.com/atreib/golangdi/entities"
)

type UserRepository struct {
}

func (*UserRepository) GetUser(login, password string) (*entities.User, error) {
	if login != "admin" || password != "132456" {
		return nil, errors.New("user does not exist")
	}

	return &entities.User{
		Id:    1,
		Login: "admin",
		Email: "admin@mail.com",
	}, nil
}
