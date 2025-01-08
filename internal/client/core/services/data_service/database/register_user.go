package database

import (
	"errors"

	"github.com/besean163/gophkeeper/internal/client/core/models"
)

func (s Service) RegisterUser(login, password string) (*models.User, error) {
	return nil, errors.New("can't register user offline")
}
