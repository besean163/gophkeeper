package services

import (
	"github.com/besean163/gophkeeper/internal/client/core/models"
)

type Repository interface{}

type DataService struct {
	repository Repository
}

func NewDataService(repository Repository) DataService {
	return DataService{
		repository: repository,
	}
}

func (s DataService) LoginUser(login, password string) (*models.User, error) {
	user := models.User{
		Login:    login,
		Password: password,
	}
	return &user, nil
	// return nil, errors.New("user not found")
}

func (s DataService) RegisterUser(login, password string) (*models.User, error) {
	user := models.User{
		Login:    login,
		Password: password,
	}
	return &user, nil
	// return nil, errors.New("user already exist")
}
