package bucket

import "github.com/besean163/gophkeeper/internal/models"

type BucketService interface {
	GetAccounts(name string) []*models.Account
}
