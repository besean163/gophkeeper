package input

import apierrors "github.com/besean163/gophkeeper/internal/server/api/errors"

// NoteUpdate структура обновления заметки.
type NoteUpdate struct {
	UUID    string `json:"uuid"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

func (i NoteUpdate) Validate(failCode int) *apierrors.Error {
	if i.UUID == "" {
		return apierrors.NewError(failCode, apierrors.ErrorEmptyUUID.Error())
	}

	return nil
}
