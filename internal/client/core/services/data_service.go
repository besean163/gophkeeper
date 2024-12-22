package services

import (
	"fmt"

	"github.com/besean163/gophkeeper/internal/client/core/models"
	"github.com/besean163/gophkeeper/internal/logger"
)

type Repository interface{}

type DataService struct {
	repository Repository
}

func NewDataService(repository Repository) DataService {
	return DataService{
		repository: repository,
	}
}

func (s DataService) LoginUser(login, password string) (*models.User, error) {
	user := &models.User{
		Login:    login,
		Password: password,
	}
	err := initUserToken(user)
	if err != nil {
		return nil, err
	}
	return user, nil
	// return nil, errors.New("user not found")
}

func (s DataService) RegisterUser(login, password string) (*models.User, error) {
	user := &models.User{
		Login:    login,
		Password: password,
	}
	err := initUserToken(user)
	if err != nil {
		return nil, err
	}
	return user, nil
	// return nil, errors.New("user already exist")
}

func initUserToken(user *models.User) error {
	// input := entities.RegisterInput{
	// 	Login:    user.Login,
	// 	Password: user.Password,
	// }

	// token, err := api.NewClient().Register(input)
	// if err != nil {
	// 	return err
	// }

	// user.Token = token.Token
	user.Token = "test token"
	return nil
}

func getAccount(id int) (*models.Account, error) {
	account := &models.Account{
		ID:       id,
		Name:     fmt.Sprintf("account_name_%d", id),
		Login:    fmt.Sprintf("account_login_%d", id),
		Password: fmt.Sprintf("account_password_%d", id),
	}
	return account, nil
}

func deleteAccount(id int) error {
	logger.Get().Printf("account deleted: %d", id)
	return nil
}

func saveAccount(account *models.Account) error {
	logger.Get().Printf("account saved: %d", account.ID)
	return nil
}
