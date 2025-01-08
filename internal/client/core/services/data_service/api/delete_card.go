package api

import (
	"github.com/besean163/gophkeeper/internal/client/core/models"
)

func (s Service) DeleteCard(user models.User, item models.Card, soft bool) error {
	var err error
	err = s.storeService.DeleteCard(user, item, soft)
	if err != nil {
		return err
	}

	err = s.syncer.SyncCards(user)
	if err != nil {
		return err
	}

	return nil
}
