package entities

import apierrors "github.com/besean163/gophkeeper/internal/server/api/errors"

// CardUpdateInput структура для обновления карты.
type CardUpdateInput struct {
	UUID   string `json:"uuid"`
	Name   string `json:"name"`
	Number int    `json:"number"`
	Exp    string `json:"exp"`
	CVV    int    `json:"cvv"`
}

func (i CardUpdateInput) Validate(failCode int) *apierrors.Error {
	if i.UUID == "" {
		return apierrors.NewError(failCode, apierrors.ErrorEmptyUUID.Error())
	}

	return nil
}
