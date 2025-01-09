package input

import apierrors "github.com/besean163/gophkeeper/internal/server/api/errors"

// NoteDelete структура для удаления заметки.
type NoteDelete struct {
	UUID string `json:"uuid"`
}

func (i NoteDelete) Validate(failCode int) *apierrors.Error {
	if i.UUID == "" {
		return apierrors.NewError(failCode, apierrors.ErrorEmptyUUID.Error())
	}

	return nil
}
