// Package auth представляет реализацию сервиса работы с данными
package bucket

import (
	models "github.com/besean163/gophkeeper/internal/models/server"

	changedetector "github.com/besean163/gophkeeper/internal/server/services/bucket/change_detector"
	"github.com/besean163/gophkeeper/internal/server/services/bucket/changes"
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
	GetAccountChanges(user models.User, compare changes.AccountCompare) changes.AccountChanges
	GetNotesChanges(user models.User, compare changes.NoteCompare) changes.NoteChanges
	GetCardsChanges(user models.User, compare changes.CardCompare) changes.CardChanges
}

type ServiceOptions struct {
	Repository
	timecontroller.TimeController
	uuidcontroller.UUIDController
	ChangeDetector
}

// Service структура сервиса
type Service struct {
	repository     Repository
	changeDetector ChangeDetector
	timeController timecontroller.TimeController
	uuidController uuidcontroller.UUIDController
}

// NewService создание структуры сервиса
func NewService(options ServiceOptions) Service {
	item := Service{
		repository:     options.Repository,
		changeDetector: options.ChangeDetector,
		timeController: options.TimeController,
		uuidController: options.UUIDController,
	}

	if item.changeDetector == nil {
		item.changeDetector = changedetector.NewChangeDetector()
	}

	return item
}
