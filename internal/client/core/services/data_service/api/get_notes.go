package api

import (
	"github.com/besean163/gophkeeper/internal/client/core/models"
)

func (s Service) GetNotes(user models.User) ([]models.Note, error) {
	return s.storeService.GetNotes(user)
}
