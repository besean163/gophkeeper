package auth

import (
	apierrors "github.com/besean163/gophkeeper/internal/server/api/errors"
)

// LoginUser авторизация пользователя
func (s Service) LoginUser(login, password string) (string, error) {
	user := s.repository.GetUserByLogin(login)

	if user == nil {
		return "", apierrors.ErrorUserNotExist
	}

	if !s.encrypter.CheckPassword(user.Password, password) {
		return "", apierrors.ErrorValidatePasswordWrong
	}

	return s.tokener.GetToken(user)
}
