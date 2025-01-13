package changedetector

import (
	clientmodels "github.com/besean163/gophkeeper/internal/models/client"
	models "github.com/besean163/gophkeeper/internal/models/server"
	"github.com/besean163/gophkeeper/internal/server/services/bucket/changes"
)

// GetNotesChanges определение изменений в заметках
func (d ChangeDetector) GetNotesChanges(user models.User, compare changes.NoteCompare) changes.NoteChanges {
	changes := changes.NewNoteChanges()

	mapExternalItems := map[string]*clientmodels.Note{}
	for _, externalItem := range compare.CompareItems {
		mapExternalItems[externalItem.UUID] = &externalItem
	}

	mapItems := map[string]*models.Note{}
	for _, item := range compare.Items {
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
			item.Content = externalItem.Content
			item.UpdatedAt = externalItem.UpdatedAt
			changes.Updated = append(changes.Updated, item)
		}
	}
	return changes
}
