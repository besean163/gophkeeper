package api

import (
	"errors"

	models "github.com/besean163/gophkeeper/internal/models/client"

	"github.com/besean163/gophkeeper/internal/server/api/entities/input"
)

func (s Service) RegisterUser(login, password string) (*models.User, error) {

	if s.apiClient.HasConnection() {
		return s.registerUserApi(login, password)
	}

	return nil, errors.New("can't register user offline")
}

func (s Service) registerUserApi(login, password string) (*models.User, error) {
	input := input.Register{
		Login:    login,
		Password: password,
	}

	output, err := s.apiClient.Register(input)
	if err != nil {
		return nil, err
	}

	token := output.Token
	encryptPassword, err := s.encrypter.Encrypt(password)
	if err != nil {
		return nil, err
	}
	user := &models.User{
		Login:    login,
		Password: encryptPassword,
		Token:    token,
	}

	// сохраняем пользователя т.к. обновился токен
	err = s.storeService.SaveUser(*user)
	if err != nil {
		return nil, err
	}

	// синхронизируем все данные пользователя на клиент
	err = s.syncer.Sync(*user, SyncNodeAccount, SyncNodeNote, SyncNodeCard)
	if err != nil {
		return nil, err
	}
	return user, nil
}
