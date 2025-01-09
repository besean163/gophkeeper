package api

import (
	models "github.com/besean163/gophkeeper/internal/models/client"
)

func (s Service) UpdateAccount(user models.User, item models.Account) error {
	var err error
	err = s.storeService.UpdateAccount(user, item)
	if err != nil {
		return err
	}
	err = s.syncer.Sync(user, SyncNodeAccount)
	if err != nil {
		return err
	}

	return nil
}
