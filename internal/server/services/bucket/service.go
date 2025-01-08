// Package auth представляет реализацию сервиса работы с данными
package bucket

import (
	"github.com/besean163/gophkeeper/internal/server/models"
	changedetector "github.com/besean163/gophkeeper/internal/server/services/bucket/change_detector"
	timecontroller "github.com/besean163/gophkeeper/internal/utils/time_controller"
	uuidcontroller "github.com/besean163/gophkeeper/internal/utils/uuid_controller"
)

// Repository интерфейс репозитория
type Repository interface {
	GetAccounts(user models.User) ([]*models.Account, error)
	GetAccount(uuid string) (*models.Account, error)
	SaveAccount(account *models.Account) error
	DeleteAccount(uuid string) error

	GetNotes(user models.User) ([]*models.Note, error)
	GetNote(uuid string) (*models.Note, error)
	SaveNote(note *models.Note) error
	DeleteNote(uuid string) error

	GetCards(user models.User) ([]*models.Card, error)
	GetCard(uuid string) (*models.Card, error)
	SaveCard(card *models.Card) error
	DeleteCard(uuid string) error
}

// ChangeDetector интерфейс работы с определителем изменений
type ChangeDetector interface {
	GetAccountChanges(user models.User, items []*models.Account, externalItems []models.ExternalAccount) (created []*models.Account, updated []*models.Account, deleted []string)
	GetNotesChanges(user models.User, items []*models.Note, externalItems []models.ExternalNote) (created []*models.Note, updated []*models.Note, deleted []string)
	GetCardsChanges(user models.User, items []*models.Card, externalItems []models.ExternalCard) (created []*models.Card, updated []*models.Card, deleted []string)
}

// Service структура сервиса
type Service struct {
	repository     Repository
	changeDetector ChangeDetector
	timeController timecontroller.TimeController
	uuidController uuidcontroller.UUIDController
}

// NewService создание структуры сервиса
func NewService(repository Repository, timeController timecontroller.TimeController, uuidController uuidcontroller.UUIDController, changeDetector ChangeDetector) Service {
	item := Service{
		repository:     repository,
		changeDetector: changeDetector,
		timeController: timeController,
		uuidController: uuidController,
	}

	if item.changeDetector == nil {
		item.changeDetector = changedetector.NewChangeDetector()
	}

	return item
}
