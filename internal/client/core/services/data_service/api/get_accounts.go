package api

import (
	"github.com/besean163/gophkeeper/internal/client/core/models"
)

func (s Service) GetAccounts(user models.User) ([]models.Account, error) {
	return s.storeService.GetAccounts(user)
}
