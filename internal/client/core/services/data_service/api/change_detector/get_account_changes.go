package changedetector

import (
	models "github.com/besean163/gophkeeper/internal/models/client"
	servermodels "github.com/besean163/gophkeeper/internal/models/server"
)

func (d ChangeDetector) GetAccountChanges(user models.User, items []models.Account, externalItems []servermodels.Account) (created []models.Account, updated []models.Account, deleted []models.Account) {
	created = make([]models.Account, 0)
	updated = make([]models.Account, 0)
	deleted = make([]models.Account, 0)

	mapItems := map[string]models.Account{}
	for _, item := range items {
		mapItems[item.UUID] = item
	}

	mapExternalItems := map[string]servermodels.Account{}
	for _, externalItem := range externalItems {
		mapExternalItems[externalItem.UUID] = externalItem
	}

	for _, item := range mapItems {
		_, ok := mapExternalItems[item.UUID]
		if !ok {
			deleted = append(deleted, item)
		}
	}

	for _, externalItem := range mapExternalItems {
		item, ok := mapItems[externalItem.UUID]
		if !ok {
			item = models.Account{
				UserID:    user.ID,
				UUID:      externalItem.UUID,
				Login:     externalItem.Login,
				Name:      externalItem.Name,
				Password:  externalItem.Password,
				CreatedAt: externalItem.CreatedAt,
				UpdatedAt: externalItem.UpdatedAt,
			}
			created = append(created, item)
		}

		if externalItem.UpdatedAt > item.UpdatedAt {
			item.Name = externalItem.Name
			item.Login = externalItem.Login
			item.Password = externalItem.Password
			item.UpdatedAt = externalItem.UpdatedAt
			updated = append(updated, item)
		}
	}

	return created, updated, deleted
}
