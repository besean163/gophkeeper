package database

import (
	models "github.com/besean163/gophkeeper/internal/models/client"
)

func (s Service) SaveUser(user models.User) error {
	exist := s.repository.GetUserByLogin(user.Login)
	if exist == nil {
		return s.repository.SaveUser(user)
	}

	exist.Token = user.Token
	return s.repository.SaveUser(*exist)
}
