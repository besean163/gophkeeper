package api

import (
	"github.com/besean163/gophkeeper/internal/client/core/models"
)

func (s Service) DeleteAccount(user models.User, item models.Account, soft bool) error {
	var err error
	err = s.storeService.DeleteAccount(user, item, soft)
	if err != nil {
		return err
	}

	err = s.syncer.SyncAccounts(user)
	if err != nil {
		return err
	}

	return nil
}
