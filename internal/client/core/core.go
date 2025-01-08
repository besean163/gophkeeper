// Package core предоставляет ядро работы клиента с сервером и локальной базой данных
package core

import (
	"github.com/besean163/gophkeeper/internal/client/core/interfaces"
	"github.com/besean163/gophkeeper/internal/client/core/models"
	defaultlogger "github.com/besean163/gophkeeper/internal/logger/default_logger"

	"github.com/besean163/gophkeeper/internal/logger"
)

// Core структура ядра
type Core struct {
	logger      logger.Logger
	dataService interfaces.DataService
	*models.User
}

// Core создание структуры ядра
func NewCore(dataService interfaces.DataService, logger logger.Logger) Core {
	if logger == nil {
		logger = defaultlogger.NewDefaultLogger()
	}
	return Core{
		logger:      logger,
		dataService: dataService,
	}
}

// Login авторизация пользователя
func (core *Core) Login(login, password string) error {

	user, err := core.dataService.LoginUser(login, password)
	if err != nil {
		return err
	}

	core.logger.Info("success login user")
	core.User = user
	return nil
}

// Register регистрация пользователя
func (core *Core) Register(login, password string) error {

	user, err := core.dataService.RegisterUser(login, password)
	if err != nil {
		return err
	}

	core.logger.Info("success register user")
	core.User = user
	return nil
}

// GetAccounts получение списка ваккаунтов
func (core *Core) GetAccounts() ([]models.Account, error) {
	return core.dataService.GetAccounts(*core.User)
}

// SaveAccount сохранение аккаунта
func (core *Core) SaveAccount(account models.Account) error {
	var err error
	if account.UUID == "" {
		err = core.dataService.CreateAccount(*core.User, account)
	} else {
		err = core.dataService.UpdateAccount(*core.User, account)
	}
	if err != nil {
		return err
	}
	return nil
}

// DeleteAccount удаление аккаунта
func (core *Core) DeleteAccount(account models.Account) error {
	err := core.dataService.DeleteAccount(*core.User, account, true)
	if err != nil {
		return err
	}
	return nil
}

// GetNotes получение списка заметок
func (core *Core) GetNotes() ([]models.Note, error) {
	return core.dataService.GetNotes(*core.User)
}

// SaveNote сохранение заметки
func (core *Core) SaveNote(item models.Note) error {
	var err error
	if item.UUID == "" {
		err = core.dataService.CreateNote(*core.User, item)
	} else {
		err = core.dataService.UpdateNote(*core.User, item)
	}
	if err != nil {
		return err
	}
	return nil
}

// DeleteNote удаление заметки
func (core *Core) DeleteNote(item models.Note) error {
	err := core.dataService.DeleteNote(*core.User, item, true)
	if err != nil {
		return err
	}
	return nil
}

// GetCards получение списка карт
func (core *Core) GetCards() ([]models.Card, error) {
	return core.dataService.GetCards(*core.User)
}

// SaveCard сохранение карты
func (core *Core) SaveCard(item models.Card) error {
	var err error
	if item.UUID == "" {
		err = core.dataService.CreateCard(*core.User, item)
	} else {
		err = core.dataService.UpdateCard(*core.User, item)
	}
	if err != nil {
		return err
	}
	return nil
}

// DeleteCard удаление карты
func (core *Core) DeleteCard(item models.Card) error {
	err := core.dataService.DeleteCard(*core.User, item, true)
	if err != nil {
		return err
	}
	return nil
}
