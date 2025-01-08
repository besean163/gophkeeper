package database

import "github.com/besean163/gophkeeper/internal/client/core/models"

func (s Service) SaveUser(user models.User) error {
	exist := s.repository.GetUserByLogin(user.Login)
	if exist == nil {
		return s.repository.SaveUser(user)
	}

	exist.Token = user.Token
	return s.repository.SaveUser(*exist)
}
