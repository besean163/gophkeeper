package changedetector

import (
	clientmodels "github.com/besean163/gophkeeper/internal/models/client"
	models "github.com/besean163/gophkeeper/internal/models/server"
	"github.com/besean163/gophkeeper/internal/server/services/bucket/changes"
)

// GetCardsChanges определение изменений в картах
func (d ChangeDetector) GetCardsChanges(user models.User, compare changes.CardCompare) changes.CardChanges {
	changes := changes.NewCardChanges()

	mapExternalItems := map[string]*clientmodels.Card{}
	for _, externalItem := range compare.CompareItems {
		mapExternalItems[externalItem.UUID] = &externalItem
	}

	mapItems := map[string]*models.Card{}
	for _, item := range compare.Items {
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
			item.Number = externalItem.Number
			item.Exp = externalItem.Exp
			item.CVV = externalItem.CVV
			item.UpdatedAt = externalItem.UpdatedAt
			changes.Updated = append(changes.Updated, item)
		}
	}
	return changes
}
