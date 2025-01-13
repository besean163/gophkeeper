package bucket

import (
	models "github.com/besean163/gophkeeper/internal/models/server"

	clientmodels "github.com/besean163/gophkeeper/internal/models/client"

	"github.com/besean163/gophkeeper/internal/server/interfaces"
	"github.com/besean163/gophkeeper/internal/server/services/bucket/changes"
)

// SyncAccounts синхронизация аккаутнов
func (s Service) SyncAccounts(service interfaces.BucketService, user models.User, externalItems []clientmodels.Account) error {
	items, err := s.repository.GetAccounts(user)
	if err != nil {
		return err
	}

	compare := changes.AccountCompare{
		Items:        items,
		CompareItems: externalItems,
	}
	changes := s.changeDetector.GetAccountChanges(user, compare)

	for _, item := range changes.Created {
		err := service.CreateAccount(user, item)
		if err != nil {
			return err
		}
	}

	for _, item := range changes.Updated {
		err := service.UpdateAccount(user, item)
		if err != nil {
			return err
		}
	}

	for _, uuid := range changes.Deleted {
		err := service.DeleteAccount(user, uuid)
		if err != nil {
			return err
		}
	}
	return nil
}
