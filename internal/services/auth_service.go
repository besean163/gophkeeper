package services

import (
	"github.com/besean163/gophkeeper/internal/models"
)

type UserRepository interface {
	GetUser(login string) (models.User, error)
}

type AuthService struct {
	repository UserRepository
}

func NewAuthService(repository UserRepository) AuthService {
	return AuthService{
		repository: repository,
	}
}

func (s AuthService) GetUser(login string) *models.User {
	return nil
}
