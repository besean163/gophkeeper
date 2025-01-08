package api

import (
	"github.com/besean163/gophkeeper/internal/client/core/models"
)

func (s Service) CreateAccount(user models.User, item models.Account) error {
	var err error
	err = s.storeService.CreateAccount(user, item)
	if err != nil {
		return err
	}

	err = s.syncer.SyncAccounts(user)
	if err != nil {
		return err
	}

	return nil
}
