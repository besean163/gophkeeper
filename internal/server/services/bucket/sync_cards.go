package bucket

import (
	clientmodels "github.com/besean163/gophkeeper/internal/models/client"
	models "github.com/besean163/gophkeeper/internal/models/server"
	"github.com/besean163/gophkeeper/internal/server/interfaces"
	"github.com/besean163/gophkeeper/internal/server/services/bucket/changes"
)

// SyncCards синхронизация карт
func (s Service) SyncCards(service interfaces.BucketService, user models.User, externalItems []clientmodels.Card) error {
	items, err := s.repository.GetCards(user)
	if err != nil {
		return err
	}

	compare := changes.CardCompare{
		Items:        items,
		CompareItems: externalItems,
	}
	changes := s.changeDetector.GetCardsChanges(user, compare)

	for _, item := range changes.Created {
		err := service.CreateCard(user, item)
		if err != nil {
			return err
		}
	}

	for _, item := range changes.Updated {
		err := service.UpdateCard(user, item)
		if err != nil {
			return err
		}
	}

	for _, uuid := range changes.Deleted {
		err := service.DeleteCard(user, uuid)
		if err != nil {
			return err
		}
	}
	return nil
}
