package changedetector

import (
	"testing"

	"github.com/besean163/gophkeeper/internal/client/core/models"
	changedetector "github.com/besean163/gophkeeper/internal/client/core/services/data_service/api/change_detector"
	"github.com/stretchr/testify/assert"
)

func TestGetNoteChanges(t *testing.T) {
	d := changedetector.NewChangeDetector()
	user := models.User{ID: 1}

	tests := []struct {
		name          string
		items         []models.Note
		externalItems []models.ExternalNote
		created       []models.Note
		updated       []models.Note
		deleted       []models.Note
	}{
		{
			name:  "created",
			items: make([]models.Note, 0),
			externalItems: []models.ExternalNote{
				{
					UUID:      "uuid_1",
					Name:      "name_1",
					Content:   "text_1",
					CreatedAt: 1,
					UpdatedAt: 1,
				},
			},
			created: []models.Note{
				{
					UserID:    user.ID,
					UUID:      "uuid_1",
					Name:      "name_1",
					Content:   "text_1",
					CreatedAt: 1,
					UpdatedAt: 1,
				},
			},
			updated: make([]models.Note, 0),
			deleted: make([]models.Note, 0),
		},
		{
			name: "not_updated",
			items: []models.Note{
				{
					UserID:    user.ID,
					UUID:      "uuid_1",
					Name:      "name_1",
					Content:   "text_1",
					CreatedAt: 1,
					UpdatedAt: 1,
				},
			},
			externalItems: []models.ExternalNote{
				{
					UUID:      "uuid_1",
					Name:      "name_1",
					Content:   "text_1",
					CreatedAt: 1,
					UpdatedAt: 1,
				},
			},
			created: make([]models.Note, 0),
			updated: make([]models.Note, 0),
			deleted: make([]models.Note, 0),
		},
		{
			name: "updated",
			items: []models.Note{
				{
					UserID:    user.ID,
					UUID:      "uuid_1",
					Name:      "name_1",
					Content:   "text_1",
					CreatedAt: 1,
					UpdatedAt: 1,
				},
			},
			externalItems: []models.ExternalNote{
				{
					UUID:      "uuid_1",
					Name:      "new_name_1",
					Content:   "text_1",
					CreatedAt: 1,
					UpdatedAt: 2,
				},
			},
			created: make([]models.Note, 0),
			updated: []models.Note{
				{
					UserID:    user.ID,
					UUID:      "uuid_1",
					Name:      "new_name_1",
					Content:   "text_1",
					CreatedAt: 1,
					UpdatedAt: 2,
				},
			},
			deleted: make([]models.Note, 0),
		},
		{
			name: "deleted",
			items: []models.Note{
				{
					UserID:    user.ID,
					UUID:      "uuid_1",
					Name:      "name_1",
					Content:   "text_1",
					CreatedAt: 1,
					UpdatedAt: 1,
				},
			},
			externalItems: []models.ExternalNote{},
			created:       make([]models.Note, 0),
			updated:       make([]models.Note, 0),
			deleted: []models.Note{
				{
					UserID:    user.ID,
					UUID:      "uuid_1",
					Name:      "name_1",
					Content:   "text_1",
					CreatedAt: 1,
					UpdatedAt: 1,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			created, updated, deleted := d.GetNoteChanges(user, test.items, test.externalItems)
			assert.Equal(t, test.created, created)
			assert.Equal(t, test.updated, updated)
			assert.Equal(t, test.deleted, deleted)
		})
	}
}
