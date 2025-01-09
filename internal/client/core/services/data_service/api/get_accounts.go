package api

import (
	models "github.com/besean163/gophkeeper/internal/models/client"
)

func (s Service) GetAccounts(user models.User) ([]models.Account, error) {
	return s.storeService.GetAccounts(user)
}
