package bucket

import (
	"github.com/besean163/gophkeeper/internal/server/interfaces"
	"github.com/besean163/gophkeeper/internal/server/models"
)

// SyncCards синхронизация карт
func (s Service) SyncCards(service interfaces.BucketService, user models.User, externalItems []models.ExternalCard) error {
	items, err := s.repository.GetCards(user)
	if err != nil {
		return err
	}

	created, updated, deleted := s.changeDetector.GetCardsChanges(user, items, externalItems)

	for _, item := range created {
		err := service.CreateCard(user, item)
		if err != nil {
			return err
		}
	}

	for _, item := range updated {
		err := service.UpdateCard(user, item)
		if err != nil {
			return err
		}
	}

	for _, uuid := range deleted {
		err := service.DeleteCard(user, uuid)
		if err != nil {
			return err
		}
	}
	return nil
}
