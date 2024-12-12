package bucket

import "github.com/besean163/gophkeeper/internal/models"

type BucketService interface {
	GetAccounts() []*models.Account
	CreateAccount(account *models.Account) error
	UpdateAccount(account *models.Account) error
	DeleteAccount(id int) error
}
