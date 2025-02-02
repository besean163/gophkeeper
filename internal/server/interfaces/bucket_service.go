package interfaces

import (
	clientmodels "github.com/besean163/gophkeeper/internal/models/client"
	models "github.com/besean163/gophkeeper/internal/models/server"
)

// BucketService интерфейс сервиса предоставления данных
type BucketService interface {
	GetAccounts(user models.User) ([]*models.Account, error)
	CreateAccount(user models.User, account *models.Account) error
	UpdateAccount(user models.User, account *models.Account) error
	DeleteAccount(user models.User, uuid string) error
	SyncAccounts(service BucketService, user models.User, accounts []clientmodels.Account) error

	GetNotes(user models.User) ([]*models.Note, error)
	CreateNote(user models.User, account *models.Note) error
	UpdateNote(user models.User, account *models.Note) error
	DeleteNote(user models.User, uuid string) error
	SyncNotes(service BucketService, user models.User, accounts []clientmodels.Note) error

	GetCards(user models.User) ([]*models.Card, error)
	CreateCard(user models.User, account *models.Card) error
	UpdateCard(user models.User, account *models.Card) error
	DeleteCard(user models.User, uuid string) error
	SyncCards(service BucketService, user models.User, accounts []clientmodels.Card) error
}
