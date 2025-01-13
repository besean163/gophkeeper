package database

import (
	models "github.com/besean163/gophkeeper/internal/models/client"
)

func (s Service) DeleteAccount(user models.User, item models.Account, soft bool) error {
	if soft {
		item.DeletedAt = s.timeController.Now()
		return s.repository.SaveAccount(item)
	}
	return s.repository.DeleteAccount(item.UUID)
}
