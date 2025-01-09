package api

import (
	models "github.com/besean163/gophkeeper/internal/models/client"
)

func (s Service) CreateAccount(user models.User, item models.Account) error {
	var err error
	err = s.storeService.CreateAccount(user, item)
	if err != nil {
		return err
	}

	err = s.syncer.Sync(user, SyncNodeAccount)
	if err != nil {
		return err
	}

	return nil
}
