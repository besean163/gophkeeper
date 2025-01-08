package entities

import apierrors "github.com/besean163/gophkeeper/internal/server/api/errors"

// AccountDeleteInput структура для удаления аккаунта.
type AccountDeleteInput struct {
	UUID string `json:"uuid"`
}

func (i AccountDeleteInput) Validate(failCode int) *apierrors.Error {
	if i.UUID == "" {
		return apierrors.NewError(failCode, apierrors.ErrorEmptyUUID.Error())
	}

	return nil
}
