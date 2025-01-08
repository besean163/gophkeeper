package changedetector

import (
	"github.com/besean163/gophkeeper/internal/client/core/models"
)

func (d ChangeDetector) GetCardChanges(user models.User, items []models.Card, externalItems []models.ExternalCard) (created []models.Card, updated []models.Card, deleted []models.Card) {
	created = make([]models.Card, 0)
	updated = make([]models.Card, 0)
	deleted = make([]models.Card, 0)

	mapItems := map[string]models.Card{}
	for _, item := range items {
		mapItems[item.UUID] = item
	}

	mapExternalItems := map[string]models.ExternalCard{}
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
			created = append(created, item)
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
