package database

import (
	models "github.com/besean163/gophkeeper/internal/models/client"

	"github.com/besean163/gophkeeper/internal/logger"
	pencrypt "github.com/besean163/gophkeeper/internal/utils/password_encrypter"
	timecontroller "github.com/besean163/gophkeeper/internal/utils/time_controller"
	uuidcontroller "github.com/besean163/gophkeeper/internal/utils/uuid_controller"
)

type Repository interface {
	GetUserByLogin(login string) *models.User
	SaveUser(user models.User) error

	GetAccounts(user models.User) ([]models.Account, error)
	SaveAccount(item models.Account) error
	DeleteAccount(uuid string) error

	GetNotes(user models.User) ([]models.Note, error)
	SaveNote(item models.Note) error
	DeleteNote(uuid string) error

	GetCards(user models.User) ([]models.Card, error)
	SaveCard(item models.Card) error
	DeleteCard(uuid string) error
}

type ServiceOptions struct {
	Repository
	logger.Logger
	pencrypt.Encrypter
	timecontroller.TimeController
	uuidcontroller.UUIDController
}

type Service struct {
	repository     Repository
	logger         logger.Logger
	encrypter      pencrypt.Encrypter
	uuidController uuidcontroller.UUIDController
	timeController timecontroller.TimeController
}

func NewService(options ServiceOptions) Service {
	return Service{
		repository:     options.Repository,
		encrypter:      options.Encrypter,
		logger:         options.Logger,
		timeController: options.TimeController,
		uuidController: options.UUIDController,
	}
}
