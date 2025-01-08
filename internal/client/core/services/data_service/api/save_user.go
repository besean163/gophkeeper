package api

import "github.com/besean163/gophkeeper/internal/client/core/models"

func (s Service) SaveUser(user models.User) error {
	return s.storeService.SaveUser(user)
}
