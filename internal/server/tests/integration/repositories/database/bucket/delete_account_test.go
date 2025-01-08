package bucket

import (
	"testing"

	"github.com/besean163/gophkeeper/internal/server/models"
	"github.com/besean163/gophkeeper/internal/server/repositories/database/bucket"
	"github.com/stretchr/testify/assert"
)

func TestDeleteAccount(t *testing.T) {
	account_1 := &models.Account{ID: 1, UUID: "00000000-0000-0000-0000-000000000001", UserID: 1, Name: "name_1", Login: "login_1", Password: "password_1", CreatedAt: 1, UpdatedAt: 1}
	account_2 := &models.Account{ID: 2, UUID: "00000000-0000-0000-0000-000000000002", UserID: 2, Name: "name_2", Login: "login_2", Password: "password_2", CreatedAt: 1, UpdatedAt: 1}

	loadFixtureAccounts(t, []*models.Account{
		account_1,
		account_2,
	})
	defer cleanUpFixtureAccounts(t)

	r := bucket.NewRepository(db)

	var err error
	var account *models.Account

	account, err = r.GetAccount("00000000-0000-0000-0000-000000000001")
	assert.Nil(t, err)
	assert.Equal(t, account_1, account)

	err = r.DeleteAccount(account.UUID)
	assert.Nil(t, err)

	account, err = r.GetAccount("00000000-0000-0000-0000-000000000001")
	assert.Nil(t, err)
	assert.Nil(t, account)
}
