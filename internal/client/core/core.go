package core

import (
	"github.com/besean163/gophkeeper/internal/client/core/api"
	"github.com/besean163/gophkeeper/internal/client/core/interfaces"
	"github.com/besean163/gophkeeper/internal/client/core/models"
	apidataService "github.com/besean163/gophkeeper/internal/client/core/services/api_data_service"
	"github.com/besean163/gophkeeper/internal/logger"
)

var instance *Core

type Core struct {
	dataService interfaces.DataService
	client      *api.Client
	*models.User
}

func Init() {
	dataService := apidataService.NewService()
	instance = &Core{
		dataService: dataService,
		client:      api.NewClient(),
	}
}

func getInstance() *Core {
	if instance == nil {
		Init()
	}
	return instance

}

func Login(login, password string) error {

	user, err := getInstance().dataService.LoginUser(login, password)
	if err != nil {
		return err
	}

	logger.Get().Println("success login user")
	getInstance().User = user
	return nil
}

func Register(login, password string) error {

	user, err := getInstance().dataService.RegisterUser(login, password)
	if err != nil {
		return err
	}

	logger.Get().Println("success register user")
	getInstance().User = user
	return nil
}

func GetAccounts() ([]models.Account, error) {
	return getInstance().dataService.GetAccounts(*getInstance().User)
}

func SaveAccount(account models.Account) error {
	err := getInstance().dataService.SaveAccount(*getInstance().User, account)
	if err != nil {
		return err
	}
	return nil
}

func DeleteAccount(account models.Account) error {
	err := getInstance().dataService.DeleteAccount(*getInstance().User, account)
	if err != nil {
		return err
	}
	return nil
}
