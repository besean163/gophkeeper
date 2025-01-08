package api

import "github.com/besean163/gophkeeper/internal/client/core/models"

func (s Service) GetUserByLogin(login string) *models.User {
	return s.storeService.GetUserByLogin(login)
}
