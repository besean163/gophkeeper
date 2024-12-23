package interfaces

import "github.com/besean163/gophkeeper/internal/client/core/models"

type DataService interface {
	LoginUser(login, password string) (*models.User, error)
	RegisterUser(login, password string) (*models.User, error)
	GetAccounts(user models.User) ([]models.Account, error)
	SaveAccount(user models.User, account models.Account) error
}
