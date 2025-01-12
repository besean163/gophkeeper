package changedetector

import (
	"github.com/besean163/gophkeeper/internal/client/core/services/data_service/api/changes"
	models "github.com/besean163/gophkeeper/internal/models/client"
	servermodels "github.com/besean163/gophkeeper/internal/models/server"
)

func (d ChangeDetector) GetAccountChanges(user models.User, compare changes.AccountCompare) changes.AccountChanges {
	changes := changes.NewAccountChanges()

	mapItems := map[string]models.Account{}
	for _, item := range compare.Items {
		mapItems[item.UUID] = item
	}

	mapExternalItems := map[string]servermodels.Account{}
	for _, externalItem := range compare.CompareItems {
		mapExternalItems[externalItem.UUID] = externalItem
	}

	for _, item := range mapItems {
		_, ok := mapExternalItems[item.UUID]
		if !ok {
			changes.Deleted = append(changes.Deleted, item)
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
			changes.Created = append(changes.Created, item)
		}

		if externalItem.UpdatedAt > item.UpdatedAt {
			item.Name = externalItem.Name
			item.Login = externalItem.Login
			item.Password = externalItem.Password
			item.UpdatedAt = externalItem.UpdatedAt
			changes.Updated = append(changes.Updated, item)
		}
	}

	return changes
}
