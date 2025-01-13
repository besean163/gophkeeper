package input

import (
	apierrors "github.com/besean163/gophkeeper/internal/server/api/errors"
	"github.com/google/uuid"
)

// NoteDelete структура для удаления заметки.
type NoteDelete struct {
	UUID string `json:"uuid"`
}

func (i NoteDelete) Validate(failCode int) *apierrors.Error {
	err := uuid.Validate(i.UUID)
	if err != nil {
		return apierrors.NewError(failCode, apierrors.ErrorEmptyUUID.Error())
	}

	return nil
}
