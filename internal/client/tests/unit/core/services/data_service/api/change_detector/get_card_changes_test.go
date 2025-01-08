package changedetector

import (
	"testing"

	"github.com/besean163/gophkeeper/internal/client/core/models"
	changedetector "github.com/besean163/gophkeeper/internal/client/core/services/data_service/api/change_detector"
	"github.com/stretchr/testify/assert"
)

func TestGetCardChanges(t *testing.T) {
	d := changedetector.NewChangeDetector()
	user := models.User{ID: 1}

	tests := []struct {
		name          string
		items         []models.Card
		externalItems []models.ExternalCard
		created       []models.Card
		updated       []models.Card
		deleted       []models.Card
	}{
		{
			name:  "created",
			items: make([]models.Card, 0),
			externalItems: []models.ExternalCard{
				{
					UUID:      "uuid_1",
					Name:      "name_1",
					Number:    1111,
					Exp:       "11|11",
					CVV:       111,
					CreatedAt: 1,
					UpdatedAt: 1,
				},
			},
			created: []models.Card{
				{
					UserID:    user.ID,
					UUID:      "uuid_1",
					Name:      "name_1",
					Number:    1111,
					Exp:       "11|11",
					CVV:       111,
					CreatedAt: 1,
					UpdatedAt: 1,
				},
			},
			updated: make([]models.Card, 0),
			deleted: make([]models.Card, 0),
		},
		{
			name: "not_updated",
			items: []models.Card{
				{
					UserID:    user.ID,
					UUID:      "uuid_1",
					Name:      "name_1",
					Number:    1111,
					Exp:       "11|11",
					CVV:       111,
					CreatedAt: 1,
					UpdatedAt: 1,
				},
			},
			externalItems: []models.ExternalCard{
				{
					UUID:      "uuid_1",
					Name:      "name_1",
					Number:    1111,
					Exp:       "11|11",
					CVV:       111,
					CreatedAt: 1,
					UpdatedAt: 1,
				},
			},
			created: make([]models.Card, 0),
			updated: make([]models.Card, 0),
			deleted: make([]models.Card, 0),
		},
		{
			name: "updated",
			items: []models.Card{
				{
					UserID:    user.ID,
					UUID:      "uuid_1",
					Name:      "name_1",
					Number:    1111,
					Exp:       "11|11",
					CVV:       111,
					CreatedAt: 1,
					UpdatedAt: 1,
				},
			},
			externalItems: []models.ExternalCard{
				{
					UUID:      "uuid_1",
					Name:      "new_name_1",
					Number:    1111,
					Exp:       "11|11",
					CVV:       111,
					CreatedAt: 1,
					UpdatedAt: 2,
				},
			},
			created: make([]models.Card, 0),
			updated: []models.Card{
				{
					UserID:    user.ID,
					UUID:      "uuid_1",
					Name:      "new_name_1",
					Number:    1111,
					Exp:       "11|11",
					CVV:       111,
					CreatedAt: 1,
					UpdatedAt: 2,
				},
			},
			deleted: make([]models.Card, 0),
		},
		{
			name: "deleted",
			items: []models.Card{
				{
					UserID:    user.ID,
					UUID:      "uuid_1",
					Name:      "name_1",
					Number:    1111,
					Exp:       "11|11",
					CVV:       111,
					CreatedAt: 1,
					UpdatedAt: 1,
				},
			},
			externalItems: []models.ExternalCard{},
			created:       make([]models.Card, 0),
			updated:       make([]models.Card, 0),
			deleted: []models.Card{
				{
					UserID:    user.ID,
					UUID:      "uuid_1",
					Name:      "name_1",
					Number:    1111,
					Exp:       "11|11",
					CVV:       111,
					CreatedAt: 1,
					UpdatedAt: 1,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			created, updated, deleted := d.GetCardChanges(user, test.items, test.externalItems)
			assert.Equal(t, test.created, created)
			assert.Equal(t, test.updated, updated)
			assert.Equal(t, test.deleted, deleted)
		})
	}
}
