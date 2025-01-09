package api

import (
	"github.com/besean163/gophkeeper/internal/logger"
	models "github.com/besean163/gophkeeper/internal/models/client"
	"github.com/besean163/gophkeeper/internal/server/api/entities/input"
)

func (s Service) syncCardsOnServer(user models.User) error {
	if !s.apiClient.HasConnection() {
		return nil
	}

	items, err := s.storeService.GetCards(user)
	if err != nil {
		return err
	}

	apiItems := make([]input.CardSync, 0)
	for _, a := range items {
		apiItem := input.CardSync{
			UUID:      a.UUID,
			Name:      a.Name,
			Number:    a.Number,
			Exp:       a.Exp,
			CVV:       a.CVV,
			CreatedAt: a.CreatedAt,
			UpdatedAt: a.UpdatedAt,
			DeletedAt: a.DeletedAt,
			SyncedAt:  a.SyncedAt,
		}
		apiItems = append(apiItems, apiItem)
	}
	input := input.CardsSync{
		Cards: apiItems,
	}

	s.apiClient.SetToken(user.Token)
	err = s.apiClient.SyncCards(input)
	if err != nil {
		return err
	}
	s.logger.Debug("sync cards on server ...", logger.Field{Key: "user", Value: user.Login})
	return nil
}
