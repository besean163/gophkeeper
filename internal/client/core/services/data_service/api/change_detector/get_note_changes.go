package changedetector

import (
	"github.com/besean163/gophkeeper/internal/client/core/services/data_service/api/changes"
	models "github.com/besean163/gophkeeper/internal/models/client"
	servermodels "github.com/besean163/gophkeeper/internal/models/server"
)

func (d ChangeDetector) GetNoteChanges(user models.User, compare changes.NoteCompare) changes.NoteChanges {
	changes := changes.NewNoteChanges()

	mapItems := map[string]models.Note{}
	for _, item := range compare.Items {
		mapItems[item.UUID] = item
	}

	mapExternalItems := map[string]servermodels.Note{}
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
			item = models.Note{
				UserID:    user.ID,
				UUID:      externalItem.UUID,
				Name:      externalItem.Name,
				Content:   externalItem.Content,
				CreatedAt: externalItem.CreatedAt,
				UpdatedAt: externalItem.UpdatedAt,
			}
			changes.Created = append(changes.Created, item)
		}

		if externalItem.UpdatedAt > item.UpdatedAt {
			item.Name = externalItem.Name
			item.Content = externalItem.Content
			item.UpdatedAt = externalItem.UpdatedAt
			changes.Updated = append(changes.Updated, item)
		}
	}

	return changes
}
