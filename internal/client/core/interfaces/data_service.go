package interfaces

import (
	models "github.com/besean163/gophkeeper/internal/models/client"
)

// DataService интерфейс сервиса для работы с данными
type DataService interface {
	GetUserByLogin(login string) *models.User
	LoginUser(login, password string) (*models.User, error)
	RegisterUser(login, password string) (*models.User, error)
	SaveUser(user models.User) error

	GetAccounts(user models.User) ([]models.Account, error)
	CreateAccount(user models.User, item models.Account) error
	UpdateAccount(user models.User, item models.Account) error
	DeleteAccount(user models.User, item models.Account, soft bool) error

	GetNotes(user models.User) ([]models.Note, error)
	CreateNote(user models.User, item models.Note) error
	UpdateNote(user models.User, item models.Note) error
	DeleteNote(user models.User, item models.Note, soft bool) error

	GetCards(user models.User) ([]models.Card, error)
	CreateCard(user models.User, item models.Card) error
	UpdateCard(user models.User, item models.Card) error
	DeleteCard(user models.User, item models.Card, soft bool) error
}
