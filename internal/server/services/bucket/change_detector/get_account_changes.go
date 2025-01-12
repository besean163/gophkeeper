package changedetector

import (
	clientmodels "github.com/besean163/gophkeeper/internal/models/client"
	models "github.com/besean163/gophkeeper/internal/models/server"
	"github.com/besean163/gophkeeper/internal/server/services/bucket/changes"
)

// GetAccountChanges определение изменений в аккаунтах
func (d ChangeDetector) GetAccountChanges(user models.User, compare changes.AccountCompare) changes.AccountChanges {
	changes := changes.NewAccountChanges()

	mapExternalItems := map[string]*clientmodels.Account{}
	for _, externalItem := range compare.CompareItems {
		mapExternalItems[externalItem.UUID] = &externalItem
	}

	mapItems := map[string]*models.Account{}
	for _, item := range compare.Items {
		mapItems[item.UUID] = item
	}

	for uuid, externalItem := range mapExternalItems {
		_, ok := mapItems[uuid]
		if !ok {
			item := &models.Account{
				UserID:    user.ID,
				UUID:      externalItem.UUID,
				Name:      externalItem.Name,
				Login:     externalItem.Login,
				Password:  externalItem.Password,
				CreatedAt: externalItem.CreatedAt,
				UpdatedAt: externalItem.UpdatedAt,
			}
			mapItems[uuid] = item
			changes.Created = append(changes.Created, item)
			continue
		}
	}

	for uuid, item := range mapItems {
		externalItem, ok := mapExternalItems[uuid]
		if !ok {
			continue
		}

		if externalItem.DeletedAt != 0 && externalItem.DeletedAt > item.UpdatedAt {
			changes.Deleted = append(changes.Deleted, item.UUID)
			continue
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
