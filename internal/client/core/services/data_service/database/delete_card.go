package database

import (
	models "github.com/besean163/gophkeeper/internal/models/client"
)

func (s Service) DeleteCard(user models.User, item models.Card, soft bool) error {
	if soft {
		item.DeletedAt = s.timeController.Now()
		return s.repository.SaveCard(item)
	}
	return s.repository.DeleteCard(item.UUID)
}
