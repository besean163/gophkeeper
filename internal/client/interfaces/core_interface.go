package interfaces

import (
	models "github.com/besean163/gophkeeper/internal/models/client"
)

type Core interface {
	Register(login, password string) error
	Login(login, password string) error
	GetAccounts() ([]models.Account, error)
	SaveAccount(item models.Account) error
	DeleteAccount(item models.Account) error

	GetNotes() ([]models.Note, error)
	SaveNote(item models.Note) error
	DeleteNote(item models.Note) error

	GetCards() ([]models.Card, error)
	SaveCard(item models.Card) error
	DeleteCard(item models.Card) error
}
