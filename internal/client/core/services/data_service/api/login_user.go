package api

import (
	"errors"

	models "github.com/besean163/gophkeeper/internal/models/client"

	"github.com/besean163/gophkeeper/internal/server/api/entities/input"
)

func (s Service) LoginUser(login, password string) (*models.User, error) {

	if s.apiClient.HasConnection() {
		return s.loginUserApi(login, password)
	}

	return s.loginUserLocally(login, password)
}

func (s Service) loginUserLocally(login, password string) (*models.User, error) {
	user, err := s.storeService.LoginUser(login, password)
	if err != nil {
		return nil, err
	}
	return user, err
}
func (s Service) loginUserApi(login, password string) (*models.User, error) {
	input := input.Login{
		Login:    login,
		Password: password,
	}

	output, err := s.apiClient.Login(input)
	if err != nil {
		return nil, err
	}

	token := output.Token
	encryptPassword, err := s.encrypter.Encrypt(password)
	if err != nil {
		return nil, err
	}
	updateUser := models.User{
		Login:    login,
		Password: encryptPassword,
		Token:    token,
	}

	// сохраняем пользователя т.к. обновился токен
	err = s.storeService.SaveUser(updateUser)
	if err != nil {
		return nil, err
	}

	user := s.storeService.GetUserByLogin(updateUser.Login)
	if user == nil {
		return nil, errors.New("something wrong")
	}

	// синхронизируем все данные пользователя на клиент
	err = s.syncer.Sync(*user, SyncNodeAccount, SyncNodeNote, SyncNodeCard)
	if err != nil {
		return nil, err
	}
	return user, nil
}
