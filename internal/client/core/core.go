package core

import (
	"github.com/besean163/gophkeeper/internal/client/core/api"
	"github.com/besean163/gophkeeper/internal/client/core/interfaces"
	"github.com/besean163/gophkeeper/internal/client/core/models"
	apidataservice "github.com/besean163/gophkeeper/internal/client/core/services/api_data_service"
	"github.com/besean163/gophkeeper/internal/logger"
)

var Instance *Core

type Core struct {
	DataService interfaces.DataService
	APIClient   api.Client
	*models.User
}

func Init() {
	if Instance != nil {
		return
	}

	// repository := database.NewRepository()
	// dataService := services.NewDataService(repository)
	dataService := apidataservice.NewService()

	Instance = &Core{
		DataService: dataService,
		APIClient:   api.NewClient(),
	}
}

func (core *Core) Login(login, password string) error {

	user, err := core.DataService.LoginUser(login, password)
	if err != nil {
		return err
	}

	logger.Get().Println("success login user")
	core.User = user
	return nil
}

func (core *Core) Register(login, password string) error {
	user, err := core.DataService.RegisterUser(login, password)
	if err != nil {
		return err
	}

	logger.Get().Println("success register user")
	core.User = user
	return nil
}

func (core *Core) GetAccounts() ([]models.Account, error) {
	// return []models.Account{
	// 	{
	// 		Name:     "account_1",
	// 		Login:    "login_1",
	// 		Password: "password_1",
	// 	},
	// 	{
	// 		Name:     "account_2",
	// 		Login:    "login_2",
	// 		Password: "password_2",
	// 	},
	// 	{
	// 		Name:     "account_3",
	// 		Login:    "login_3",
	// 		Password: "password_3",
	// 	},
	// }

	return core.DataService.GetAccounts(*core.User)
}

func (core *Core) Save(account models.Account) error {
	return nil
}
