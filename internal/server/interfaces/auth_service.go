package interfaces

import "github.com/besean163/gophkeeper/internal/server/models"

type AuthService interface {
	GetUser(login string) *models.User
	RegisterUser(user *models.User) (string, error)
	CreateUserToken(user *models.User) (string, error)
}
