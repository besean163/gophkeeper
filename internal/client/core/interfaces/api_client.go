package interfaces

import "github.com/besean163/gophkeeper/internal/server/api/entities"

// ApiClient интерфейс сервиса для работы с сервером
type ApiClient interface {
	HasConnection() bool
	Register(input entities.RegisterInput) (entities.TokenOutput, error)
	Login(input entities.LoginInput) (entities.TokenOutput, error)
	SetToken(token string)

	SyncAccounts(input entities.AccountsSyncInput) error
	GetAccounts() (*entities.GetAccountsOutput, error)

	SyncNotes(input entities.NotesSyncInput) error
	GetNotes() (*entities.GetNotesOutput, error)

	SyncCards(input entities.CardsSyncInput) error
	GetCards() (*entities.GetCardsOutput, error)
}
