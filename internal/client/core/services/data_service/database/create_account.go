package database

import (
	models "github.com/besean163/gophkeeper/internal/models/client"
)

func (s Service) CreateAccount(user models.User, item models.Account) error {
	item.UserID = user.ID
	item.CreatedAt = s.timeController.Now()
	item.UpdatedAt = s.timeController.Now()

	return s.repository.SaveAccount(item)
}
