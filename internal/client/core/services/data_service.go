package services

import (
	"github.com/besean163/gophkeeper/internal/client/core/models"
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
