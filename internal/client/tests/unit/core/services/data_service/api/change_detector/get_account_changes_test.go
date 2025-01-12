package changedetector

import (
	"testing"

	models "github.com/besean163/gophkeeper/internal/models/client"
	servermodels "github.com/besean163/gophkeeper/internal/models/server"

	changedetector "github.com/besean163/gophkeeper/internal/client/core/services/data_service/api/change_detector"
	"github.com/besean163/gophkeeper/internal/client/core/services/data_service/api/changes"
	"github.com/stretchr/testify/assert"
)

func TestGetAccountChanges(t *testing.T) {
	d := changedetector.NewChangeDetector()
	user := models.User{ID: 1}

	tests := []struct {
		name          string
		items         []models.Account
		externalItems []servermodels.Account
		created       []models.Account
		updated       []models.Account
		deleted       []models.Account
	}{
		{
			name:  "created",
			items: make([]models.Account, 0),
			externalItems: []servermodels.Account{
				{
					UUID:      "uuid_1",
					Name:      "name_1",
					Login:     "login_1",
					Password:  "password_1",
					CreatedAt: 1,
					UpdatedAt: 1,
				},
			},
			created: []models.Account{
				{
					UserID:    user.ID,
					UUID:      "uuid_1",
					Name:      "name_1",
					Login:     "login_1",
					Password:  "password_1",
					CreatedAt: 1,
					UpdatedAt: 1,
				},
			},
			updated: make([]models.Account, 0),
			deleted: make([]models.Account, 0),
		},
		{
			name: "not_updated",
			items: []models.Account{
				{
					UserID:    user.ID,
					UUID:      "uuid_1",
					Name:      "name_1",
					Login:     "login_1",
					Password:  "password_1",
					CreatedAt: 1,
					UpdatedAt: 1,
				},
			},
			externalItems: []servermodels.Account{
				{
					UUID:      "uuid_1",
					Name:      "name_1",
					Login:     "login_1",
					Password:  "password_1",
					CreatedAt: 1,
					UpdatedAt: 1,
				},
			},
			created: make([]models.Account, 0),
			updated: make([]models.Account, 0),
			deleted: make([]models.Account, 0),
		},
		{
			name: "updated",
			items: []models.Account{
				{
					UserID:    user.ID,
					UUID:      "uuid_1",
					Name:      "name_1",
					Login:     "login_1",
					Password:  "password_1",
					CreatedAt: 1,
					UpdatedAt: 1,
				},
			},
			externalItems: []servermodels.Account{
				{
					UUID:      "uuid_1",
					Name:      "new_name_1",
					Login:     "new_login_1",
					Password:  "new_password_1",
					CreatedAt: 1,
					UpdatedAt: 2,
				},
			},
			created: make([]models.Account, 0),
			updated: []models.Account{
				{
					UserID:    user.ID,
					UUID:      "uuid_1",
					Name:      "new_name_1",
					Login:     "new_login_1",
					Password:  "new_password_1",
					CreatedAt: 1,
					UpdatedAt: 2,
				},
			},
			deleted: make([]models.Account, 0),
		},
		{
			name: "deleted",
			items: []models.Account{
				{
					UserID:    user.ID,
					UUID:      "uuid_1",
					Name:      "name_1",
					Login:     "login_1",
					Password:  "password_1",
					CreatedAt: 1,
					UpdatedAt: 1,
				},
			},
			externalItems: []servermodels.Account{},
			created:       make([]models.Account, 0),
			updated:       make([]models.Account, 0),
			deleted: []models.Account{
				{
					UserID:    user.ID,
					UUID:      "uuid_1",
					Name:      "name_1",
					Login:     "login_1",
					Password:  "password_1",
					CreatedAt: 1,
					UpdatedAt: 1,
				},
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			compare := changes.AccountCompare{
				Items:        test.items,
				CompareItems: test.externalItems,
			}
			changes := d.GetAccountChanges(user, compare)
			assert.Equal(t, test.created, changes.Created)
			assert.Equal(t, test.updated, changes.Updated)
			assert.Equal(t, test.deleted, changes.Deleted)
		})
	}
}
