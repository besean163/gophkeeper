package entities

import apierrors "github.com/besean163/gophkeeper/internal/server/api/errors"

// LoginInput структура для авторизации пользователя.
type LoginInput struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (i LoginInput) Validate(failCode int) *apierrors.Error {
	if i.Login == "" {
		return apierrors.NewError(failCode, apierrors.ErrorValidateLoginEmpty.Error())
	}

	if i.Password == "" {
		return apierrors.NewError(failCode, apierrors.ErrorValidatePasswordEmpty.Error())
	}

	return nil
}
