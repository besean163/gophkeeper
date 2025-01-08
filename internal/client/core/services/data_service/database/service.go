package database

import (
	"github.com/besean163/gophkeeper/internal/client/core/models"
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

type Service struct {
	logger         logger.Logger
	encrypter      pencrypt.Encrypter
	repository     Repository
	uuidController uuidcontroller.UUIDController
	timeController timecontroller.TimeController
}

func NewService(repository Repository, encrypter pencrypt.Encrypter, logger logger.Logger, timeController timecontroller.TimeController, uuidController uuidcontroller.UUIDController) Service {
	return Service{
		repository:     repository,
		encrypter:      encrypter,
		logger:         logger,
		timeController: timeController,
		uuidController: uuidController,
	}
}
