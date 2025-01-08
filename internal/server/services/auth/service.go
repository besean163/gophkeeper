// Package auth представляет реализацию сервиса работы с пользователями
package auth

import (
	"github.com/besean163/gophkeeper/internal/server/models"
	apitoken "github.com/besean163/gophkeeper/internal/utils/api_token"
	pencrypt "github.com/besean163/gophkeeper/internal/utils/password_encrypter"
	timecontroller "github.com/besean163/gophkeeper/internal/utils/time_controller"
	uuidcontroller "github.com/besean163/gophkeeper/internal/utils/uuid_controller"
)

// Repository интерфейс репозитория
type Repository interface {
	GetUser(id int) (*models.User, error)
	GetUserByLogin(login string) *models.User
	SaveUser(user *models.User) error
}

// Service структура сервиса
type Service struct {
	repository     Repository
	encrypter      pencrypt.Encrypter
	tokener        apitoken.Tokener
	timeController timecontroller.TimeController
	uuidController uuidcontroller.UUIDController
}

// NewService создание структуры сервиса
func NewService(repository Repository, encrypter pencrypt.Encrypter, tokener apitoken.Tokener, timeController timecontroller.TimeController, uuidController uuidcontroller.UUIDController) Service {
	return Service{
		repository:     repository,
		encrypter:      encrypter,
		tokener:        tokener,
		timeController: timeController,
		uuidController: uuidController,
	}
}
