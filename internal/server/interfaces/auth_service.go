package interfaces

import "github.com/besean163/gophkeeper/internal/server/models"

// AuthService интерфейс сервиса авторизации
type AuthService interface {
	GetUser(id int) (*models.User, error)
	RegisterUser(login, password string) (string, error)
	LoginUser(login, password string) (string, error)
}
