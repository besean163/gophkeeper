package api

import (
	"github.com/besean163/gophkeeper/internal/logger"
	models "github.com/besean163/gophkeeper/internal/models/client"
	"github.com/besean163/gophkeeper/internal/server/api/entities/input"
)

func (s Service) syncAccountsOnServer(user models.User) error {
	if !s.apiClient.HasConnection() {
		return nil
	}

	items, err := s.storeService.GetAccounts(user)
	if err != nil {
		return err
	}

	apiItems := make([]input.AccountSync, 0)
	for _, a := range items {
		apiItem := input.AccountSync{
			UUID:      a.UUID,
			Name:      a.Name,
			Login:     a.Login,
			Password:  a.Password,
			CreatedAt: a.CreatedAt,
			UpdatedAt: a.UpdatedAt,
			DeletedAt: a.DeletedAt,
			SyncedAt:  a.SyncedAt,
		}
		apiItems = append(apiItems, apiItem)
	}
	input := input.AccountsSync{
		Accounts: apiItems,
	}

	s.apiClient.SetToken(user.Token)
	err = s.apiClient.SyncAccounts(input)
	if err != nil {
		return err
	}
	s.logger.Debug("sync accounts on server ...", logger.Field{Key: "user", Value: user.Login})
	return nil
}
