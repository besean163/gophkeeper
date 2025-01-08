package changedetector

import "github.com/besean163/gophkeeper/internal/server/models"

// GetAccountChanges определение изменений в аккаунтах
func (d ChangeDetector) GetAccountChanges(user models.User, items []*models.Account, externalItems []models.ExternalAccount) (created []*models.Account, updated []*models.Account, deleted []string) {
	created = make([]*models.Account, 0)
	updated = make([]*models.Account, 0)
	deleted = make([]string, 0)

	mapExternalItems := map[string]*models.ExternalAccount{}
	for _, externalItem := range externalItems {
		mapExternalItems[externalItem.UUID] = &externalItem
	}

	mapItems := map[string]*models.Account{}
	for _, item := range items {
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
			created = append(created, item)
			continue
		}
	}

	for uuid, item := range mapItems {
		externalItem, ok := mapExternalItems[uuid]
		if !ok {
			continue
		}

		if externalItem.DeletedAt != 0 && externalItem.DeletedAt > item.UpdatedAt {
			deleted = append(deleted, item.UUID)
			continue
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
