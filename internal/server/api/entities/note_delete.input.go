package entities

import apierrors "github.com/besean163/gophkeeper/internal/server/api/errors"

// NoteDeleteInput структура для удаления заметки.
type NoteDeleteInput struct {
	UUID string `json:"uuid"`
}

func (i NoteDeleteInput) Validate(failCode int) *apierrors.Error {
	if i.UUID == "" {
		return apierrors.NewError(failCode, apierrors.ErrorEmptyUUID.Error())
	}

	return nil
}
