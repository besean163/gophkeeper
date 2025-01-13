package database

import (
	"errors"

	models "github.com/besean163/gophkeeper/internal/models/client"
)

func (s Service) RegisterUser(login, password string) (*models.User, error) {
	return nil, errors.New("can't register user offline")
}
