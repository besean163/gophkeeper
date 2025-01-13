package database

import (
	models "github.com/besean163/gophkeeper/internal/models/client"
)

func (s Service) UpdateAccount(user models.User, item models.Account) error {
	item.UserID = user.ID
	item.UpdatedAt = s.timeController.Now()
	return s.repository.SaveAccount(item)
}
