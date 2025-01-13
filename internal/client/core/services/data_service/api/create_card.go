package api

import (
	models "github.com/besean163/gophkeeper/internal/models/client"
)

func (s Service) CreateCard(user models.User, item models.Card) error {
	var err error
	err = s.storeService.CreateCard(user, item)
	if err != nil {
		return err
	}

	err = s.syncer.Sync(user, SyncNodeCard)
	if err != nil {
		return err
	}

	return nil
}
