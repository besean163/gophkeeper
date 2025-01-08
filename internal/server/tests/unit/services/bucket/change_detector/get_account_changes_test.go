package changedetector

import (
	"testing"

	"github.com/besean163/gophkeeper/internal/server/models"
	changedetector "github.com/besean163/gophkeeper/internal/server/services/bucket/change_detector"
	"github.com/stretchr/testify/assert"
)

func TestGetAccountChanges(t *testing.T) {
	d := changedetector.NewChangeDetector()
	user := models.User{ID: 1}

	tests := []struct {
		name             string
		accounts         []*models.Account
		externalAccounts []models.ExternalAccount
		created          []*models.Account
		updated          []*models.Account
		deleted          []string
	}{
		{
			name:     "created",
			accounts: make([]*models.Account, 0),
			externalAccounts: []models.ExternalAccount{
				{
					UUID:      "uuid_1",
					Name:      "name_1",
					Login:     "login_1",
					Password:  "passwrod_1",
					CreatedAt: 1,
					UpdatedAt: 1,
					DeletedAt: 0,
					SyncedAt:  0,
				},
			},
			created: []*models.Account{
				{
					UserID:    user.ID,
					UUID:      "uuid_1",
					Name:      "name_1",
					Login:     "login_1",
					Password:  "passwrod_1",
					CreatedAt: 1,
					UpdatedAt: 1,
				},
			},
			updated: make([]*models.Account, 0),
			deleted: make([]string, 0),
		},
		{
			name: "not_updated",
			accounts: []*models.Account{
				{
					UserID:    user.ID,
					UUID:      "uuid_1",
					Name:      "name_1",
					Login:     "login_1",
					Password:  "passwrod_1",
					CreatedAt: 1,
					UpdatedAt: 1,
				},
			},
			externalAccounts: []models.ExternalAccount{
				{
					UUID:      "uuid_1",
					Name:      "name_1",
					Login:     "login_1",
					Password:  "passwrod_1",
					CreatedAt: 1,
					UpdatedAt: 1,
					DeletedAt: 0,
					SyncedAt:  0,
				},
			},
			created: make([]*models.Account, 0),
			updated: make([]*models.Account, 0),
			deleted: make([]string, 0),
		},
		{
			name: "updated",
			accounts: []*models.Account{
				{
					UserID:    user.ID,
					UUID:      "uuid_1",
					Name:      "name_1",
					Login:     "login_1",
					Password:  "passwrod_1",
					CreatedAt: 1,
					UpdatedAt: 1,
				},
			},
			externalAccounts: []models.ExternalAccount{
				{
					UUID:      "uuid_1",
					Name:      "new_name_1",
					Login:     "new_login_1",
					Password:  "new_passwrod_1",
					CreatedAt: 1,
					UpdatedAt: 2,
					DeletedAt: 0,
					SyncedAt:  0,
				},
			},
			created: make([]*models.Account, 0),
			updated: []*models.Account{
				{
					UserID:    user.ID,
					UUID:      "uuid_1",
					Name:      "new_name_1",
					Login:     "new_login_1",
					Password:  "new_passwrod_1",
					CreatedAt: 1,
					UpdatedAt: 2,
				},
			},
			deleted: make([]string, 0),
		},
		{
			name: "deleted",
			accounts: []*models.Account{
				{
					UserID:    user.ID,
					UUID:      "uuid_1",
					Name:      "name_1",
					Login:     "login_1",
					Password:  "passwrod_1",
					CreatedAt: 1,
					UpdatedAt: 1,
				},
			},
			externalAccounts: []models.ExternalAccount{
				{
					UUID:      "uuid_1",
					Name:      "new_name_1",
					Login:     "new_login_1",
					Password:  "new_passwrod_1",
					CreatedAt: 1,
					UpdatedAt: 2,
					DeletedAt: 3,
					SyncedAt:  0,
				},
			},
			created: make([]*models.Account, 0),
			updated: make([]*models.Account, 0),
			deleted: []string{"uuid_1"},
		},
		{
			name: "ignore",
			accounts: []*models.Account{
				{
					UserID:    user.ID,
					UUID:      "uuid_1",
					Name:      "name_1",
					Login:     "login_1",
					Password:  "passwrod_1",
					CreatedAt: 1,
					UpdatedAt: 1,
				},
			},
			externalAccounts: make([]models.ExternalAccount, 0),
			created:          make([]*models.Account, 0),
			updated:          make([]*models.Account, 0),
			deleted:          make([]string, 0),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			created, updated, deleted := d.GetAccountChanges(user, test.accounts, test.externalAccounts)
			assert.Equal(t, test.created, created)
			assert.Equal(t, test.updated, updated)
			assert.Equal(t, test.deleted, deleted)
		})
	}
}
