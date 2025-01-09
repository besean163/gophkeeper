package database

import (
	models "github.com/besean163/gophkeeper/internal/models/client"
)

func (s Service) GetAccounts(user models.User) ([]models.Account, error) {
	return s.repository.GetAccounts(user)
}
