package database

import (
	models "github.com/besean163/gophkeeper/internal/models/client"
)

// DeleteNote удаление заметки
func (r Repository) DeleteNote(uuid string) error {
	result := r.DB.Where("uuid = ?", uuid).Delete(&models.Note{})
	err := result.Error
	if err != nil {
		return err
	}
	return nil
}
