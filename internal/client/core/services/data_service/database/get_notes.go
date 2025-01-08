package database

import "github.com/besean163/gophkeeper/internal/client/core/models"

func (s Service) GetNotes(user models.User) ([]models.Note, error) {
	return s.repository.GetNotes(user)
}
