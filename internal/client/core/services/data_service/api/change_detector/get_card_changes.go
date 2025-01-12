package changedetector

import (
	"github.com/besean163/gophkeeper/internal/client/core/services/data_service/api/changes"
	models "github.com/besean163/gophkeeper/internal/models/client"
	servermodels "github.com/besean163/gophkeeper/internal/models/server"
)

func (d ChangeDetector) GetCardChanges(user models.User, compare changes.CardCompare) changes.CardChanges {
	changes := changes.NewCardChanges()

	mapItems := map[string]models.Card{}
	for _, item := range compare.Items {
		mapItems[item.UUID] = item
	}

	mapExternalItems := map[string]servermodels.Card{}
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
			item = models.Card{
				UserID:    user.ID,
				UUID:      externalItem.UUID,
				Name:      externalItem.Name,
				Number:    externalItem.Number,
				Exp:       externalItem.Exp,
				CVV:       externalItem.CVV,
				CreatedAt: externalItem.CreatedAt,
				UpdatedAt: externalItem.UpdatedAt,
			}
			changes.Created = append(changes.Created, item)
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
