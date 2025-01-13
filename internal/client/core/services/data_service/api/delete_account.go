package api

import (
	models "github.com/besean163/gophkeeper/internal/models/client"
)

func (s Service) DeleteAccount(user models.User, item models.Account, soft bool) error {
	var err error
	err = s.storeService.DeleteAccount(user, item, soft)
	if err != nil {
		return err
	}

	err = s.syncer.Sync(user, SyncNodeAccount)
	if err != nil {
		return err
	}

	return nil
}
