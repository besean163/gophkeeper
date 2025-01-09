package interfaces

import (
	models "github.com/besean163/gophkeeper/internal/models/server"
)

// AuthService интерфейс сервиса авторизации
type AuthService interface {
	GetUser(id int) (*models.User, error)
	RegisterUser(login, password string) (string, error)
	LoginUser(login, password string) (string, error)
}
