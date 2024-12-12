package services

import "github.com/besean163/gophkeeper/internal/models"

type BucketRepository interface {
}

type BucketService struct {
	repository BucketRepository
}

func NewBucketService(repository BucketRepository) BucketService {
	return BucketService{
		repository: repository,
	}
}

func (s BucketService) GetAccounts() []*models.Account {
	accounts := make([]*models.Account, 0)
	// accounts = append(accounts, &models.Account{
	// 	Name: "test",
	// })
	return accounts
}

func (s BucketService) CreateAccount(account *models.Account) error {
	return nil
}

func (s BucketService) UpdateAccount(account *models.Account) error {
	return nil
}

func (s BucketService) DeleteAccount(id int) error {
	return nil
}
