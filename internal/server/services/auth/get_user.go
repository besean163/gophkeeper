package auth

import "github.com/besean163/gophkeeper/internal/server/models"

// GetUser получение пользователя
func (s Service) GetUser(id int) (*models.User, error) {
	return s.repository.GetUser(id)
}
