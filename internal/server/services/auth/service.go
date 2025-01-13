// Package auth представляет реализацию сервиса работы с пользователями
package auth

import (
	models "github.com/besean163/gophkeeper/internal/models/server"
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

type ServiceOptions struct {
	Repository
	pencrypt.Encrypter
	apitoken.Tokener
	timecontroller.TimeController
	uuidcontroller.UUIDController
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
func NewService(options ServiceOptions) Service {
	return Service{
		repository:     options.Repository,
		encrypter:      options.Encrypter,
		tokener:        options.Tokener,
		timeController: options.TimeController,
		uuidController: options.UUIDController,
	}
}
