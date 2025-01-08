package api

import (
	"github.com/besean163/gophkeeper/internal/client/core/models"
)

func (s Service) UpdateAccount(user models.User, item models.Account) error {
	var err error
	err = s.storeService.UpdateAccount(user, item)
	if err != nil {
		return err
	}
	err = s.syncer.SyncAccounts(user)
	if err != nil {
		return err
	}

	return nil
}
