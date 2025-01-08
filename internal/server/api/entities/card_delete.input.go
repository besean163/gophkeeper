package entities

import apierrors "github.com/besean163/gophkeeper/internal/server/api/errors"

// CardDeleteInput структура удаления карты.
type CardDeleteInput struct {
	UUID string `json:"uuid"`
}

func (i CardDeleteInput) Validate(failCode int) *apierrors.Error {
	if i.UUID == "" {
		return apierrors.NewError(failCode, apierrors.ErrorEmptyUUID.Error())
	}

	return nil
}
