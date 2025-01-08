package database

import "github.com/besean163/gophkeeper/internal/client/core/models"

func (s Service) GetCards(user models.User) ([]models.Card, error) {
	return s.repository.GetCards(user)
}
