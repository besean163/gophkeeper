package api

import (
	models "github.com/besean163/gophkeeper/internal/models/client"
)

func (s Service) SaveUser(user models.User) error {
	return s.storeService.SaveUser(user)
}
