package route

import apierrors "github.com/besean163/gophkeeper/internal/errors/api_errors"

type input struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (i input) validate(failCode int) *apierrors.Error {
	if i.Login == "" {
		return makeError(failCode, ErrorValidateLoginEmpty.Error())
	}

	if i.Password == "" {
		return makeError(failCode, ErrorValidatePasswordEmpty.Error())
	}

	return nil
}

type output struct {
	Token string `json:"token"`
}
