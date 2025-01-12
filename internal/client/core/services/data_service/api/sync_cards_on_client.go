package api

import (
	"github.com/besean163/gophkeeper/internal/client/core/services/data_service/api/changes"
	"github.com/besean163/gophkeeper/internal/logger"
	models "github.com/besean163/gophkeeper/internal/models/client"
	servermodels "github.com/besean163/gophkeeper/internal/models/server"
)

func (s Service) syncCardsOnClient(user models.User) error {
	if !s.apiClient.HasConnection() {
		return nil
	}

	items, err := s.storeService.GetCards(user)
	if err != nil {
		return err
	}

	s.apiClient.SetToken(user.Token)
	apiOutput, err := s.apiClient.GetCards()
	if err != nil {
		return err
	}

	externalItems := make([]servermodels.Card, 0)
	for _, apiItem := range apiOutput.Cards {
		externalItem := servermodels.Card{
			UUID:      apiItem.UUID,
			UserID:    user.ID,
			Name:      apiItem.Name,
			Number:    apiItem.Number,
			Exp:       apiItem.Exp,
			CVV:       apiItem.CVV,
			CreatedAt: apiItem.CreatedAt,
			UpdatedAt: apiItem.UpdatedAt,
		}
		externalItems = append(externalItems, externalItem)
	}

	compare := changes.CardCompare{
		Items:        items,
		CompareItems: externalItems,
	}
	changes := s.changeDetector.GetCardChanges(user, compare)

	for _, item := range changes.Created {
		err := s.storeService.CreateCard(user, item)
		if err != nil {
			return err
		}
	}

	for _, item := range changes.Updated {
		err := s.storeService.UpdateCard(user, item)
		if err != nil {
			return err
		}
	}

	for _, item := range changes.Deleted {
		err := s.storeService.DeleteCard(user, item, false)
		if err != nil {
			return err
		}
	}

	s.logger.Debug("sync cards on client ...", logger.Field{Key: "user", Value: user.Login})
	return nil
}
