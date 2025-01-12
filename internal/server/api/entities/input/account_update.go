package input

import (
	apierrors "github.com/besean163/gophkeeper/internal/server/api/errors"
	"github.com/google/uuid"
)

// AccountUpdate структура для обновления аккаунта.
type AccountUpdate struct {
	UUID     string `json:"uuid"`
	Name     string `json:"name"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

func (i AccountUpdate) Validate(failCode int) *apierrors.Error {
	err := uuid.Validate(i.UUID)
	if err != nil {
		return apierrors.NewError(failCode, apierrors.ErrorEmptyUUID.Error())
	}

	return nil
}
