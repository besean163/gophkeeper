package interfaces

import (
	"github.com/besean163/gophkeeper/internal/server/api/entities/input"
	"github.com/besean163/gophkeeper/internal/server/api/entities/output"
)

// ApiClient интерфейс сервиса для работы с сервером
type ApiClient interface {
	HasConnection() bool
	Register(input input.Register) (output.Token, error)
	Login(input input.Login) (output.Token, error)
	SetToken(token string)

	SyncAccounts(input input.AccountsSync) error
	GetAccounts() (*output.GetAccounts, error)

	SyncNotes(input input.NotesSync) error
	GetNotes() (*output.GetNotes, error)

	SyncCards(input input.CardsSync) error
	GetCards() (*output.GetCards, error)
}
