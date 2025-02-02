package api

import (
	"github.com/besean163/gophkeeper/internal/client/core/services/data_service/api/changes"
	"github.com/besean163/gophkeeper/internal/logger"
	models "github.com/besean163/gophkeeper/internal/models/client"
	servermodels "github.com/besean163/gophkeeper/internal/models/server"
)

func (s Service) syncAccountsOnClient(user models.User) error {
	if !s.apiClient.HasConnection() {
		return nil
	}

	items, err := s.storeService.GetAccounts(user)
	if err != nil {
		return err
	}

	s.apiClient.SetToken(user.Token)
	apiOutput, err := s.apiClient.GetAccounts()
	if err != nil {
		return err
	}

	externalItems := make([]servermodels.Account, 0)
	for _, apiItem := range apiOutput.Accounts {
		externalItem := servermodels.Account{
			UUID:      apiItem.UUID,
			UserID:    user.ID,
			Name:      apiItem.Name,
			Login:     apiItem.Login,
			Password:  apiItem.Password,
			CreatedAt: apiItem.CreatedAt,
			UpdatedAt: apiItem.UpdatedAt,
		}
		externalItems = append(externalItems, externalItem)
	}

	compare := changes.AccountCompare{
		Items:        items,
		CompareItems: externalItems,
	}
	changes := s.changeDetector.GetAccountChanges(user, compare)

	for _, item := range changes.Created {
		err := s.storeService.CreateAccount(user, item)
		if err != nil {
			return err
		}
	}

	for _, item := range changes.Updated {
		err := s.storeService.UpdateAccount(user, item)
		if err != nil {
			return err
		}
	}

	for _, item := range changes.Deleted {
		err := s.storeService.DeleteAccount(user, item, false)
		if err != nil {
			return err
		}
	}

	s.logger.Debug("sync accounts on client ...", logger.Field{Key: "user", Value: user.Login})
	return nil
}
