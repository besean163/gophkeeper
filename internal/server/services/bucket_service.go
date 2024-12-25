package services

import (
	"errors"
	"time"

	"github.com/besean163/gophkeeper/internal/server/models"
)

type BucketRepository interface {
	GetAccount(id int) (*models.Account, error)
	GetAccounts(user models.User) ([]*models.Account, error)
	SaveAccount(account models.Account) error
	DeleteAccount(id int) error
}

type BucketService struct {
	repository BucketRepository
}

func NewBucketService(repository BucketRepository) BucketService {
	return BucketService{
		repository: repository,
	}
}

func (s BucketService) GetAccounts(user models.User) ([]*models.Account, error) {
	items, err := s.repository.GetAccounts(user)
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (s BucketService) CreateAccount(account *models.Account) error {
	return s.repository.SaveAccount(*account)
}

func (s BucketService) UpdateAccount(account *models.Account) error {
	exist, err := s.repository.GetAccount(account.ID)
	if err != nil {
		return err
	}

	if exist == nil {
		return errors.New("account not found")
	}

	account.CreatedAt = exist.CreatedAt
	account.UpdatedAt = time.Now()

	return s.repository.SaveAccount(*account)
}

func (s BucketService) DeleteAccount(id int) error {
	return s.repository.DeleteAccount(id)
}
