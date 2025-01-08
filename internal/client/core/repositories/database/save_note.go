package database

import "github.com/besean163/gophkeeper/internal/client/core/models"

// SaveNote сохранение заметки
func (r Repository) SaveNote(item models.Note) error {
	if item.ID == 0 {
		return r.insertItem(&item)
	}

	return r.updateItem(&item)
}
