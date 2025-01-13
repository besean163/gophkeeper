package api

import models "github.com/besean163/gophkeeper/internal/models/client"

func (s Service) SyncAccounts(user models.User) error {
	var err error
	err = s.syncAccountsOnServer(user)
	if err != nil {
		return err
	}
	err = s.syncAccountsOnClient(user)
	if err != nil {
		return err
	}
	return nil
}
