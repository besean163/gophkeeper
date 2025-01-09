package changedetector

import (
	clientmodels "github.com/besean163/gophkeeper/internal/models/client"
	models "github.com/besean163/gophkeeper/internal/models/server"
)

// GetNotesChanges определение изменений в заметках
func (d ChangeDetector) GetNotesChanges(user models.User, items []*models.Note, externalItems []clientmodels.Note) (created []*models.Note, updated []*models.Note, deleted []string) {
	created = make([]*models.Note, 0)
	updated = make([]*models.Note, 0)
	deleted = make([]string, 0)

	mapExternalItems := map[string]*clientmodels.Note{}
	for _, externalItem := range externalItems {
		mapExternalItems[externalItem.UUID] = &externalItem
	}

	mapItems := map[string]*models.Note{}
	for _, item := range items {
		mapItems[item.UUID] = item
	}

	for uuid, externalItem := range mapExternalItems {
		_, ok := mapItems[uuid]
		if !ok {
			item := &models.Note{
				UserID:    user.ID,
				UUID:      externalItem.UUID,
				Name:      externalItem.Name,
				Content:   externalItem.Content,
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
			item.Content = externalItem.Content
			item.UpdatedAt = externalItem.UpdatedAt
			updated = append(updated, item)
		}
	}
	return created, updated, deleted
}
