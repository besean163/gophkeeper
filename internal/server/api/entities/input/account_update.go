package input

import apierrors "github.com/besean163/gophkeeper/internal/server/api/errors"

// AccountUpdate структура для обновления аккаунта.
type AccountUpdate struct {
	UUID     string `json:"uuid"`
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (i AccountUpdate) Validate(failCode int) *apierrors.Error {
	if i.UUID == "" {
		return apierrors.NewError(failCode, apierrors.ErrorEmptyUUID.Error())
	}

	return nil
}
