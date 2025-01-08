package database

import (
	"errors"

	"github.com/besean163/gophkeeper/internal/client/core/models"
)

func (s Service) LoginUser(login, password string) (*models.User, error) {
	user := s.repository.GetUserByLogin(login)

	if user == nil {
		return nil, errors.New("user not exist")
	}

	if !s.encrypter.CheckPassword(user.Password, password) {
		return nil, errors.New("wrong password")
	}

	return user, nil
}
