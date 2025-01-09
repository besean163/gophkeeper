package input

import apierrors "github.com/besean163/gophkeeper/internal/server/api/errors"

// CardDelete структура удаления карты.
type CardDelete struct {
	UUID string `json:"uuid"`
}

func (i CardDelete) Validate(failCode int) *apierrors.Error {
	if i.UUID == "" {
		return apierrors.NewError(failCode, apierrors.ErrorEmptyUUID.Error())
	}

	return nil
}
