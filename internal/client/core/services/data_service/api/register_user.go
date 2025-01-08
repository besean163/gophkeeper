package api

import (
	"github.com/besean163/gophkeeper/internal/client/core/models"
	"github.com/besean163/gophkeeper/internal/server/api/entities"
)

func (s Service) RegisterUser(login, password string) (*models.User, error) {
	var user *models.User
	var err error

	if s.apiClient.HasConnection() {
		input := entities.RegisterInput{
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
		user = &models.User{
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
		err = s.syncer.SyncAll(*user)
		if err != nil {
			return nil, err
		}
	} else {
		user, err = s.storeService.RegisterUser(login, password)
		if err != nil {
			return nil, err
		}
	}

	return user, nil
}
