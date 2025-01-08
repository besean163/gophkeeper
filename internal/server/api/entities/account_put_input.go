package entities

import apierrors "github.com/besean163/gophkeeper/internal/server/api/errors"

// AccountPutInput структура для обновления аккаунта.
type AccountPutInput struct {
	UUID     string `json:"uuid"`
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (i AccountPutInput) Validate(failCode int) *apierrors.Error {
	if i.UUID == "" {
		return apierrors.NewError(failCode, apierrors.ErrorEmptyUUID.Error())
	}

	return nil
}
