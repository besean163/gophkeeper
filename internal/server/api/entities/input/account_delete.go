package input

import apierrors "github.com/besean163/gophkeeper/internal/server/api/errors"

// AccountDeleteInput структура для удаления аккаунта.
type AccountDelete struct {
	UUID string `json:"uuid"`
}

func (i AccountDelete) Validate(failCode int) *apierrors.Error {
	if i.UUID == "" {
		return apierrors.NewError(failCode, apierrors.ErrorEmptyUUID.Error())
	}

	return nil
}
