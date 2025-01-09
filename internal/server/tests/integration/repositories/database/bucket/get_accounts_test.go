package bucket

import (
	"testing"

	models "github.com/besean163/gophkeeper/internal/models/server"
	"github.com/besean163/gophkeeper/internal/server/repositories/database/bucket"
	utilmock "github.com/besean163/gophkeeper/internal/tests/mocks/utils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetAccounts(t *testing.T) {
	user := models.User{ID: 1}
	account_1 := &models.Account{UUID: "00000000-0000-0000-0000-000000000001", UserID: 1, Name: "name_1", Login: "login_1", Password: "password_1", CreatedAt: 1, UpdatedAt: 1}
	account_2 := &models.Account{UUID: "00000000-0000-0000-0000-000000000002", UserID: 2, Name: "name_2", Login: "login_2", Password: "password_2", CreatedAt: 1, UpdatedAt: 1}
	account_3 := &models.Account{UUID: "00000000-0000-0000-0000-000000000003", UserID: 1, Name: "name_3", Login: "login_3", Password: "password_3", CreatedAt: 1, UpdatedAt: 1}

	loadFixtureAccounts(t, []*models.Account{
		account_1,
		account_2,
		account_3,
	})
	defer cleanUpFixtureAccounts(t)

	ctrl := gomock.NewController(t)
	uuidController := utilmock.NewMockUUIDController(ctrl)
	r := bucket.NewRepository(db, uuidController)

	var accounts []*models.Account
	var err error

	accounts, err = r.GetAccounts(user)
	assert.Nil(t, err)
	assert.Equal(t, []*models.Account{
		account_1,
		account_3,
	}, accounts)
}
