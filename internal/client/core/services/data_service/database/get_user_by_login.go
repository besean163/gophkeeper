package database

import (
	models "github.com/besean163/gophkeeper/internal/models/client"
)

func (s Service) GetUserByLogin(login string) *models.User {
	return s.repository.GetUserByLogin(login)
}
