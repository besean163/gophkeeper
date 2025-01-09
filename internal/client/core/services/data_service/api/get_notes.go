package api

import (
	models "github.com/besean163/gophkeeper/internal/models/client"
)

func (s Service) GetNotes(user models.User) ([]models.Note, error) {
	return s.storeService.GetNotes(user)
}
