package bucket

import (
	"testing"

	"github.com/besean163/gophkeeper/internal/server/models"
	"github.com/besean163/gophkeeper/internal/server/repositories/database/bucket"
	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	user := models.User{ID: 1}
	account_1 := &models.Account{ID: 1, UUID: "00000000-0000-0000-0000-000000000001", UserID: 1, Name: "name_1", Login: "login_1", Password: "password_1", CreatedAt: 1, UpdatedAt: 1}
	account_2 := &models.Account{ID: 2, UUID: "00000000-0000-0000-0000-000000000002", UserID: 2, Name: "name_2", Login: "login_2", Password: "password_2", CreatedAt: 1, UpdatedAt: 1}

	loadFixtureAccounts(t, []*models.Account{
		account_1,
		account_2,
	})
	defer cleanUpFixtureAccounts(t)

	r := bucket.NewRepository(db)

	var err error
	var account models.Account
	account = models.Account{ID: 3, Name: "name_3", Login: "login_3", Password: "password_3", CreatedAt: 1, UpdatedAt: 1}

	err = r.SaveAccount(&account)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "empty uuid")

	account = models.Account{ID: 3, UUID: "00000000-0000-0000-0000-000000000003", Name: "name_3", Login: "login_3", Password: "password_3", CreatedAt: 1, UpdatedAt: 1}
	err = r.SaveAccount(&account)
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "empty user id")

	account = models.Account{ID: 3, UserID: user.ID, UUID: "00000000-0000-0000-0000-000000000003", Name: "name_3", Login: "login_3", Password: "password_3", CreatedAt: 1, UpdatedAt: 1}
	err = r.SaveAccount(&account)
	assert.Nil(t, err)

	var createAccount *models.Account
	createAccount, err = r.GetAccount("00000000-0000-0000-0000-000000000003")
	assert.Nil(t, err)
	assert.Equal(t, &models.Account{ID: 3, UserID: user.ID, UUID: "00000000-0000-0000-0000-000000000003", Name: "name_3", Login: "login_3", Password: "password_3", CreatedAt: 1, UpdatedAt: 1}, createAccount)
}
