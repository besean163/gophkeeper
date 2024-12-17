package core

import (
	"github.com/besean163/gophkeeper/internal/client/core/api"
	"github.com/besean163/gophkeeper/internal/client/core/api/entities"
	"github.com/besean163/gophkeeper/internal/client/core/interfaces"
	"github.com/besean163/gophkeeper/internal/client/core/models"
	"github.com/besean163/gophkeeper/internal/client/core/repositories/database"
	"github.com/besean163/gophkeeper/internal/client/core/services"
	"github.com/besean163/gophkeeper/internal/client/tui/logger"
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

	repository := database.NewRepository()
	dataService := services.NewDataService(repository)

	Instance = &Core{
		DataService: dataService,
		APIClient:   api.NewClient(),
	}
}

func (core *Core) GetToken(login, password string) (string, error) {
	input := entities.RegisterInput{
		Login:    login,
		Password: password,
	}

	token, err := core.APIClient.Register(input)
	if err != nil {
		return "", err
	}

	return token.Token, nil
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
