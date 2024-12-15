package repositories

import (
	"github.com/besean163/gophkeeper/internal/server/models"
)

type UserRepository struct {
}

func NewUserRepository() UserRepository {
	return UserRepository{}
}

func (r UserRepository) GetUser(login string) (models.User, error) {
	return models.User{
		Login: login,
	}, nil
	// return models.User{}, errors.New("user not found")
}
