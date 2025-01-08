package bucket

import (
	"github.com/besean163/gophkeeper/internal/server/interfaces"
	"github.com/besean163/gophkeeper/internal/server/models"
)

// SyncAccounts синхронизация аккаутнов
func (s Service) SyncAccounts(service interfaces.BucketService, user models.User, externalItems []models.ExternalAccount) error {
	items, err := s.repository.GetAccounts(user)
	if err != nil {
		return err
	}

	created, updated, deleted := s.changeDetector.GetAccountChanges(user, items, externalItems)

	for _, item := range created {
		err := service.CreateAccount(user, item)
		if err != nil {
			return err
		}
	}

	for _, item := range updated {
		err := service.UpdateAccount(user, item)
		if err != nil {
			return err
		}
	}

	for _, uuid := range deleted {
		err := service.DeleteAccount(user, uuid)
		if err != nil {
			return err
		}
	}
	return nil
}
