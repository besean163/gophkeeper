package bucket

import models "github.com/besean163/gophkeeper/internal/models/server"

// DeleteNote удаление заметки
func (s Service) DeleteNote(user models.User, uuid string) error {
	return s.repository.DeleteNote(uuid)
}
