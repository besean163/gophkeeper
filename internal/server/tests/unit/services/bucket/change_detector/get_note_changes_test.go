package changedetector

import (
	"testing"

	clientmodels "github.com/besean163/gophkeeper/internal/models/client"
	models "github.com/besean163/gophkeeper/internal/models/server"
	changedetector "github.com/besean163/gophkeeper/internal/server/services/bucket/change_detector"
	"github.com/besean163/gophkeeper/internal/server/services/bucket/changes"
	"github.com/stretchr/testify/assert"
)

func TestGetNoteChanges(t *testing.T) {
	d := changedetector.NewChangeDetector()
	user := models.User{ID: 1}

	tests := []struct {
		name          string
		items         []*models.Note
		externalItems []clientmodels.Note
		created       []*models.Note
		updated       []*models.Note
		deleted       []string
	}{
		{
			name:  "created",
			items: make([]*models.Note, 0),
			externalItems: []clientmodels.Note{
				{
					UUID:      "uuid_1",
					Name:      "name_1",
					Content:   "text_1",
					CreatedAt: 1,
					UpdatedAt: 1,
					DeletedAt: 0,
					SyncedAt:  0,
				},
			},
			created: []*models.Note{
				{
					UserID:    user.ID,
					UUID:      "uuid_1",
					Name:      "name_1",
					Content:   "text_1",
					CreatedAt: 1,
					UpdatedAt: 1,
				},
			},
			updated: make([]*models.Note, 0),
			deleted: make([]string, 0),
		},
		{
			name: "not_updated",
			items: []*models.Note{
				{
					UserID:    user.ID,
					UUID:      "uuid_1",
					Name:      "name_1",
					Content:   "text_1",
					CreatedAt: 1,
					UpdatedAt: 1,
				},
			},
			externalItems: []clientmodels.Note{
				{
					UUID:      "uuid_1",
					Name:      "name_1",
					Content:   "text_1",
					CreatedAt: 1,
					UpdatedAt: 1,
					DeletedAt: 0,
					SyncedAt:  0,
				},
			},
			created: make([]*models.Note, 0),
			updated: make([]*models.Note, 0),
			deleted: make([]string, 0),
		},
		{
			name: "updated",
			items: []*models.Note{
				{
					UserID:    user.ID,
					UUID:      "uuid_1",
					Name:      "name_1",
					Content:   "text_1",
					CreatedAt: 1,
					UpdatedAt: 1,
				},
			},
			externalItems: []clientmodels.Note{
				{
					UUID:      "uuid_1",
					Name:      "new_name_1",
					Content:   "text_1",
					CreatedAt: 1,
					UpdatedAt: 2,
					DeletedAt: 0,
					SyncedAt:  0,
				},
			},
			created: make([]*models.Note, 0),
			updated: []*models.Note{
				{
					UserID:    user.ID,
					UUID:      "uuid_1",
					Name:      "new_name_1",
					Content:   "text_1",
					CreatedAt: 1,
					UpdatedAt: 2,
				},
			},
			deleted: make([]string, 0),
		},
		{
			name: "deleted",
			items: []*models.Note{
				{
					UserID:    user.ID,
					UUID:      "uuid_1",
					Name:      "name_1",
					Content:   "text_1",
					CreatedAt: 1,
					UpdatedAt: 1,
				},
			},
			externalItems: []clientmodels.Note{
				{
					UUID:      "uuid_1",
					Name:      "new_name_1",
					Content:   "text_1",
					CreatedAt: 1,
					UpdatedAt: 2,
					DeletedAt: 3,
					SyncedAt:  0,
				},
			},
			created: make([]*models.Note, 0),
			updated: make([]*models.Note, 0),
			deleted: []string{"uuid_1"},
		},
		{
			name: "ignore",
			items: []*models.Note{
				{
					UserID:    user.ID,
					UUID:      "uuid_1",
					Name:      "name_1",
					Content:   "text_1",
					CreatedAt: 1,
					UpdatedAt: 1,
				},
			},
			externalItems: make([]clientmodels.Note, 0),
			created:       make([]*models.Note, 0),
			updated:       make([]*models.Note, 0),
			deleted:       make([]string, 0),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			compare := changes.NoteCompare{
				Items:        test.items,
				CompareItems: test.externalItems,
			}
			changes := d.GetNotesChanges(user, compare)
			assert.Equal(t, test.created, changes.Created)
			assert.Equal(t, test.updated, changes.Updated)
			assert.Equal(t, test.deleted, changes.Deleted)
		})
	}
}
