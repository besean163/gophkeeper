package bucket

import (
	"errors"

	"github.com/besean163/gophkeeper/internal/server/models"
)

// SaveNote сохранение заметки
func (r Repository) SaveNote(item *models.Note) error {
	if item.UUID == "" {
		return errors.New("empty uuid")
	}

	if item.UserID == 0 {
		return errors.New("empty user id")
	}

	if item.ID == 0 {
		return r.insertItem(item)
	}

	return r.updateItem(item)
}
