package entities

import apierrors "github.com/besean163/gophkeeper/internal/server/api/errors"

// NoteUpdateInput структура обновления заметки.
type NoteUpdateInput struct {
	UUID    string `json:"uuid"`
	Name    string `json:"name"`
	Content string `json:"content"`
}

func (i NoteUpdateInput) Validate(failCode int) *apierrors.Error {
	if i.UUID == "" {
		return apierrors.NewError(failCode, apierrors.ErrorEmptyUUID.Error())
	}

	return nil
}
