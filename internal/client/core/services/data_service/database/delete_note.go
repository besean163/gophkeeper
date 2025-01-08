package database

import "github.com/besean163/gophkeeper/internal/client/core/models"

func (s Service) DeleteNote(user models.User, item models.Note, soft bool) error {
	if soft {
		item.DeletedAt = s.timeController.Now()
		return s.repository.SaveNote(item)
	}
	return s.repository.DeleteNote(item.UUID)
}
