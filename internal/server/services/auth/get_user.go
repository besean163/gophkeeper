package auth

import models "github.com/besean163/gophkeeper/internal/models/server"

// GetUser получение пользователя
func (s Service) GetUser(id int) (*models.User, error) {
	return s.repository.GetUser(id)
}
