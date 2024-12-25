package interfaces

import "github.com/besean163/gophkeeper/internal/server/models"

type BucketService interface {
	GetAccounts(user models.User) ([]*models.Account, error)
	CreateAccount(account *models.Account) error
	UpdateAccount(account *models.Account) error
	DeleteAccount(id int) error
}
