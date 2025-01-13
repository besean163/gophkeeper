package bucket

import models "github.com/besean163/gophkeeper/internal/models/server"

// DeleteNote удаление заметки
func (r Repository) DeleteNote(uuid string) error {
	result := r.db.Where("uuid = ?", uuid).Delete(&models.Note{})

	err := result.Error
	if err != nil {
		return err
	}
	return nil
}
