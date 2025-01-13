package database

import (
	models "github.com/besean163/gophkeeper/internal/models/client"
)

func (s Service) GetNotes(user models.User) ([]models.Note, error) {
	return s.repository.GetNotes(user)
}
