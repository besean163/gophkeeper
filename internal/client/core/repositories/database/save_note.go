package database

import (
	models "github.com/besean163/gophkeeper/internal/models/client"
)

// SaveNote сохранение заметки
func (r Repository) SaveNote(item models.Note) error {
	if item.UUID == "" {
		item.UUID = r.UUIDController.GetUUID()
		return r.createItem(&item)
	}

	return r.updateItem(&item)
}
