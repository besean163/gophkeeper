package api

import (
	models "github.com/besean163/gophkeeper/internal/models/client"
)

func (s Service) DeleteCard(user models.User, item models.Card, soft bool) error {
	var err error
	err = s.storeService.DeleteCard(user, item, soft)
	if err != nil {
		return err
	}

	err = s.syncer.Sync(user, SyncNodeCard)
	if err != nil {
		return err
	}

	return nil
}
