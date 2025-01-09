package database

import (
	models "github.com/besean163/gophkeeper/internal/models/client"
)

func (s Service) GetCards(user models.User) ([]models.Card, error) {
	return s.repository.GetCards(user)
}
