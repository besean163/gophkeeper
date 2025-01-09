package changedetector

import (
	models "github.com/besean163/gophkeeper/internal/models/client"
	servermodels "github.com/besean163/gophkeeper/internal/models/server"
)

func (d ChangeDetector) GetNoteChanges(user models.User, items []models.Note, externalItems []servermodels.Note) (created []models.Note, updated []models.Note, deleted []models.Note) {
	created = make([]models.Note, 0)
	updated = make([]models.Note, 0)
	deleted = make([]models.Note, 0)

	mapItems := map[string]models.Note{}
	for _, item := range items {
		mapItems[item.UUID] = item
	}

	mapExternalItems := map[string]servermodels.Note{}
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
			item = models.Note{
				UserID:    user.ID,
				UUID:      externalItem.UUID,
				Name:      externalItem.Name,
				Content:   externalItem.Content,
				CreatedAt: externalItem.CreatedAt,
				UpdatedAt: externalItem.UpdatedAt,
			}
			created = append(created, item)
		}

		if externalItem.UpdatedAt > item.UpdatedAt {
			item.Name = externalItem.Name
			item.Content = externalItem.Content
			item.UpdatedAt = externalItem.UpdatedAt
			updated = append(updated, item)
		}
	}

	return created, updated, deleted
}
