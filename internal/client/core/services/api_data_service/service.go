package apidataservice

import (
	"github.com/besean163/gophkeeper/internal/client/core/api"
	"github.com/besean163/gophkeeper/internal/client/core/api/entities"
	"github.com/besean163/gophkeeper/internal/client/core/models"
	"github.com/besean163/gophkeeper/internal/logger"
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
	logger.Debug("user with token", user)

	return user, nil
}

func (s Service) RegisterUser(login, password string) (*models.User, error) {
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

func (s Service) GetAccounts(user models.User) ([]models.Account, error) {
	output, err := api.NewClient().GetAccounts(user.Token)
	if err != nil {
		return nil, err
	}

	accounts := make([]models.Account, 0)
	for _, item := range output.Accounts {
		account := models.Account{
			Name:      item.Name,
			Login:     item.Login,
			Password:  item.Password,
			CreatedAt: item.CreatedAt,
		}
		accounts = append(accounts, account)
	}

	return accounts, nil
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

// func getAccount(id int) (*models.Account, error) {
// 	account := &models.Account{
// 		ID:       id,
// 		Name:     fmt.Sprintf("account_name_%d", id),
// 		Login:    fmt.Sprintf("account_login_%d", id),
// 		Password: fmt.Sprintf("account_password_%d", id),
// 	}
// 	return account, nil
// }

// func deleteAccount(id int) error {
// 	logger.Get().Printf("account deleted: %d", id)
// 	return nil
// }

// func saveAccount(account *models.Account) error {
// 	logger.Get().Printf("account saved: %d", account.ID)
// 	return nil
// }
