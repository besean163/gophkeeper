package bucket

import (
	"testing"

	models "github.com/besean163/gophkeeper/internal/models/server"
	utilmock "github.com/besean163/gophkeeper/internal/tests/mocks/utils"

	"github.com/besean163/gophkeeper/internal/server/repositories/database/bucket"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateAccount(t *testing.T) {
	user := models.User{ID: 1}
	item_1 := &models.Account{UUID: "00000000-0000-0000-0000-000000000001", UserID: 1, Name: "name_1", Login: "login_1", Password: "password_1", CreatedAt: 1, UpdatedAt: 1}
	item_2 := &models.Account{UUID: "00000000-0000-0000-0000-000000000002", UserID: 2, Name: "name_2", Login: "login_2", Password: "password_2", CreatedAt: 1, UpdatedAt: 1}

	loadFixtureAccounts(t, []*models.Account{
		item_1,
		item_2,
	})
	defer cleanUpFixtureAccounts(t)

	ctrl := gomock.NewController(t)
	uuidController := utilmock.NewMockUUIDController(ctrl)
	uuidController.EXPECT().GetUUID().Return("00000000-0000-0000-0000-000000000003").Times(1)
	r := bucket.NewRepository(db, uuidController)

	var err error

	item := models.Account{UserID: user.ID, Name: "name_3", Login: "login_3", Password: "password_3", CreatedAt: 1, UpdatedAt: 1}
	err = r.SaveAccount(&item)
	assert.Nil(t, err)

	var createItem *models.Account
	createItem, err = r.GetAccount("00000000-0000-0000-0000-000000000003")
	assert.Nil(t, err)
	assert.Equal(t, &models.Account{UserID: user.ID, UUID: "00000000-0000-0000-0000-000000000003", Name: "name_3", Login: "login_3", Password: "password_3", CreatedAt: 1, UpdatedAt: 1}, createItem)
}
