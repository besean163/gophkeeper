package auth

import (
	apierrors "github.com/besean163/gophkeeper/internal/server/api/errors"
	"github.com/besean163/gophkeeper/internal/server/models"
)

// RegisterUser регистрация пользователя
func (s Service) RegisterUser(login, password string) (string, error) {
	var user *models.User
	user = s.repository.GetUserByLogin(login)

	if user != nil {
		return "", apierrors.ErrorUserExist
	}

	encryptPassword, err := s.encrypter.Encrypt(password)
	if err != nil {
		return "", err
	}

	user = &models.User{
		UUID:      s.uuidController.GetUUID(),
		Login:     login,
		Password:  encryptPassword,
		CreatedAt: s.timeController.Now(),
	}

	err = s.repository.SaveUser(user)
	if err != nil {
		return "", err
	}

	return s.tokener.GetToken(user)
}
