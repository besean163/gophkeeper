package changedetector

import (
	"testing"

	clientmodels "github.com/besean163/gophkeeper/internal/models/client"

	models "github.com/besean163/gophkeeper/internal/models/server"
	changedetector "github.com/besean163/gophkeeper/internal/server/services/bucket/change_detector"
	"github.com/besean163/gophkeeper/internal/server/services/bucket/changes"
	"github.com/stretchr/testify/assert"
)

func TestGetCardChanges(t *testing.T) {
	d := changedetector.NewChangeDetector()
	user := models.User{ID: 1}

	tests := []struct {
		name          string
		items         []*models.Card
		externalItems []clientmodels.Card
		created       []*models.Card
		updated       []*models.Card
		deleted       []string
	}{
		{
			name:  "created",
			items: make([]*models.Card, 0),
			externalItems: []clientmodels.Card{
				{
					UUID:      "uuid_1",
					Name:      "name_1",
					Number:    1111,
					Exp:       "11|11",
					CVV:       111,
					CreatedAt: 1,
					UpdatedAt: 1,
					DeletedAt: 0,
					SyncedAt:  0,
				},
			},
			created: []*models.Card{
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
			updated: make([]*models.Card, 0),
			deleted: make([]string, 0),
		},
		{
			name: "not_updated",
			items: []*models.Card{
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
			externalItems: []clientmodels.Card{
				{
					UUID:      "uuid_1",
					Name:      "name_1",
					Number:    1111,
					Exp:       "11|11",
					CVV:       111,
					CreatedAt: 1,
					UpdatedAt: 1,
					DeletedAt: 0,
					SyncedAt:  0,
				},
			},
			created: make([]*models.Card, 0),
			updated: make([]*models.Card, 0),
			deleted: make([]string, 0),
		},
		{
			name: "updated",
			items: []*models.Card{
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
			externalItems: []clientmodels.Card{
				{
					UUID:      "uuid_1",
					Name:      "new_name_1",
					Number:    1111,
					Exp:       "11|11",
					CVV:       111,
					CreatedAt: 1,
					UpdatedAt: 2,
					DeletedAt: 0,
					SyncedAt:  0,
				},
			},
			created: make([]*models.Card, 0),
			updated: []*models.Card{
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
			deleted: make([]string, 0),
		},
		{
			name: "deleted",
			items: []*models.Card{
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
			externalItems: []clientmodels.Card{
				{
					UUID:      "uuid_1",
					Name:      "new_name_1",
					Number:    1111,
					Exp:       "11|11",
					CVV:       111,
					CreatedAt: 1,
					UpdatedAt: 2,
					DeletedAt: 3,
					SyncedAt:  0,
				},
			},
			created: make([]*models.Card, 0),
			updated: make([]*models.Card, 0),
			deleted: []string{"uuid_1"},
		},
		{
			name: "ignore",
			items: []*models.Card{
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
			externalItems: make([]clientmodels.Card, 0),
			created:       make([]*models.Card, 0),
			updated:       make([]*models.Card, 0),
			deleted:       make([]string, 0),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			compare := changes.CardCompare{
				Items:        test.items,
				CompareItems: test.externalItems,
			}
			changes := d.GetCardsChanges(user, compare)
			assert.Equal(t, test.created, changes.Created)
			assert.Equal(t, test.updated, changes.Updated)
			assert.Equal(t, test.deleted, changes.Deleted)
		})
	}
}
