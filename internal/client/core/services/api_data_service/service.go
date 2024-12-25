package apidataservice

import (
	"github.com/besean163/gophkeeper/internal/client/core/api"
	"github.com/besean163/gophkeeper/internal/client/core/api/entities"
	"github.com/besean163/gophkeeper/internal/client/core/models"
)

type Service struct {
}

func NewService() Service {
	return Service{}
}

func (s Service) LoginUser(login, password string) (*models.User, error) {
	user := &models.User{
		Login:    login,
		Password: password,
	}

	input := entities.GetTokenInput{
		Login:    user.Login,
		Password: user.Password,
	}

	token, err := api.NewClient().Login(input)
	if err != nil {
		return nil, err
	}
	user.Token = token.Token

	return user, nil
}

func (s Service) RegisterUser(login, password string) (*models.User, error) {
	user := &models.User{
		Login:    login,
		Password: password,
	}

	input := entities.GetTokenInput{
		Login:    user.Login,
		Password: user.Password,
	}

	token, err := api.NewClient().Register(input)
	if err != nil {
		return nil, err
	}
	user.Token = token.Token

	return user, nil
}

func (s Service) GetAccounts(user models.User) ([]models.Account, error) {
	output, err := api.NewClient().SetToken(user.Token).GetAccounts()
	if err != nil {
		return nil, err
	}

	accounts := make([]models.Account, 0)
	for _, item := range output.Accounts {
		account := models.Account{
			ID:        item.ID,
			Name:      item.Name,
			Login:     item.Login,
			Password:  item.Password,
			CreatedAt: item.CreatedAt,
		}
		accounts = append(accounts, account)
	}

	return accounts, nil
}

func (s Service) SaveAccount(user models.User, account models.Account) error {
	input := entities.AccountInput{
		ID:       account.ID,
		Name:     account.Name,
		Login:    account.Login,
		Password: account.Password,
	}

	err := api.NewClient().SetToken(user.Token).SaveAccount(input)
	if err != nil {
		return err
	}

	return nil
}

func (s Service) DeleteAccount(user models.User, account models.Account) error {
	input := entities.AccountDeleteInput{
		ID: account.ID,
	}

	err := api.NewClient().SetToken(user.Token).DeleteAccount(input)
	if err != nil {
		return err
	}

	return nil
}
