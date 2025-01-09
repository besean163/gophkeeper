package changedetector

import (
	clientmodels "github.com/besean163/gophkeeper/internal/models/client"
	models "github.com/besean163/gophkeeper/internal/models/server"
)

// GetCardsChanges определение изменений в картах
func (d ChangeDetector) GetCardsChanges(user models.User, items []*models.Card, externalItems []clientmodels.Card) (created []*models.Card, updated []*models.Card, deleted []string) {
	created = make([]*models.Card, 0)
	updated = make([]*models.Card, 0)
	deleted = make([]string, 0)

	mapExternalItems := map[string]*clientmodels.Card{}
	for _, externalItem := range externalItems {
		mapExternalItems[externalItem.UUID] = &externalItem
	}

	mapItems := map[string]*models.Card{}
	for _, item := range items {
		mapItems[item.UUID] = item
	}

	for uuid, externalItem := range mapExternalItems {
		_, ok := mapItems[uuid]
		if !ok {
			item := &models.Card{
				UserID:    user.ID,
				UUID:      externalItem.UUID,
				Name:      externalItem.Name,
				Number:    externalItem.Number,
				Exp:       externalItem.Exp,
				CVV:       externalItem.CVV,
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
			item.Number = externalItem.Number
			item.Exp = externalItem.Exp
			item.CVV = externalItem.CVV
			item.UpdatedAt = externalItem.UpdatedAt
			updated = append(updated, item)
		}
	}
	return created, updated, deleted
}
