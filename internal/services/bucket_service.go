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

func (s BucketService) GetAccounts(name string) []*models.Account {
	return nil
}
