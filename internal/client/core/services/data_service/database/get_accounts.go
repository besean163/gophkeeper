package database

import "github.com/besean163/gophkeeper/internal/client/core/models"

func (s Service) GetAccounts(user models.User) ([]models.Account, error) {
	return s.repository.GetAccounts(user)
}
