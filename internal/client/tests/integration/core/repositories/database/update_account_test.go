package database

import (
	"testing"

	models "github.com/besean163/gophkeeper/internal/models/client"
	utilmock "github.com/besean163/gophkeeper/internal/tests/mocks/utils"
	"github.com/golang/mock/gomock"

	"github.com/besean163/gophkeeper/internal/client/core/repositories/database"
	"github.com/stretchr/testify/assert"
)

func TestUpdateAccount(t *testing.T) {
	user := &models.User{ID: 1}
	item_1 := &models.Account{UUID: "00000000-0000-0000-0000-000000000001", UserID: 1, Name: "name_1", Login: "login_1", Password: "password_1"}

	loadFixtureAccounts(t, []*models.Account{
		item_1,
	})
	defer cleanUpFixtureAccounts(t)

	ctrl := gomock.NewController(t)
	uuidController := utilmock.NewMockUUIDController(ctrl)
	r := database.NewRepository(db, uuidController)
	var err error

	items, err := r.GetAccounts(*user)
	assert.Nil(t, err)
	assert.Equal(t, []models.Account{
		*item_1,
	}, items)

	item_change := &models.Account{UUID: "00000000-0000-0000-0000-000000000001", UserID: 1, Name: "new_name", Login: "new_login", Password: "password_1"}
	err = r.SaveAccount(*item_change)
	assert.Nil(t, err)

	items, err = r.GetAccounts(*user)
	assert.Nil(t, err)
	assert.Equal(t, []models.Account{
		*item_change,
	}, items)
}
