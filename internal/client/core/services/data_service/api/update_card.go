package api

import (
	"github.com/besean163/gophkeeper/internal/client/core/models"
)

func (s Service) UpdateCard(user models.User, item models.Card) error {
	var err error
	err = s.storeService.UpdateCard(user, item)
	if err != nil {
		return err
	}
	err = s.syncer.SyncCards(user)
	if err != nil {
		return err
	}

	return nil
}
