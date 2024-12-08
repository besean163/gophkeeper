package repository

import "github.com/besean163/gophkeeper/internal/models"

type UserRepository struct {
}

func NewUserRepository() UserRepository {
	return UserRepository{}
}

func (r UserRepository) GetUser(login string) (models.User, error) {
	return models.User{}, nil
}
