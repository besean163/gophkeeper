package database

import (
	models "github.com/besean163/gophkeeper/internal/models/client"
)

func (s Service) UpdateCard(user models.User, item models.Card) error {
	item.UserID = user.ID
	item.UpdatedAt = s.timeController.Now()

	return s.repository.SaveCard(item)
}
