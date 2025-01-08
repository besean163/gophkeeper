package bucket

import "github.com/besean163/gophkeeper/internal/server/models"

// DeleteNote удаление заметки
func (s Service) DeleteNote(user models.User, uuid string) error {
	return s.repository.DeleteNote(uuid)
}
