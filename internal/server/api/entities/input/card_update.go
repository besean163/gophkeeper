package input

import (
	apierrors "github.com/besean163/gophkeeper/internal/server/api/errors"
	"github.com/google/uuid"
)

// CardUpdate структура для обновления карты.
type CardUpdate struct {
	UUID   string `json:"uuid"`
	Name   string `json:"name"`
	Number int    `json:"number"`
	Exp    string `json:"exp"`
	CVV    int    `json:"cvv"`
}

func (i CardUpdate) Validate(failCode int) *apierrors.Error {
	err := uuid.Validate(i.UUID)
	if err != nil {
		return apierrors.NewError(failCode, apierrors.ErrorEmptyUUID.Error())
	}

	return nil
}
