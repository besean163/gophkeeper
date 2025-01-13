package database

import (
	"testing"

	models "github.com/besean163/gophkeeper/internal/models/client"
	utilmock "github.com/besean163/gophkeeper/internal/tests/mocks/utils"
	"github.com/golang/mock/gomock"

	"github.com/besean163/gophkeeper/internal/client/core/repositories/database"
	"github.com/stretchr/testify/assert"
)

func TestUpdateUser(t *testing.T) {
	user_1 := &models.User{ID: 1, Login: "login"}
	loadFixtureUsers(t, []*models.User{
		user_1,
	})
	defer cleanUpFixtureUser(t)

	ctrl := gomock.NewController(t)
	uuidController := utilmock.NewMockUUIDController(ctrl)
	r := database.NewRepository(db, uuidController)
	var user *models.User

	user = r.GetUserByLogin("login_new")
	assert.Nil(t, user)

	user = r.GetUserByLogin("login")
	assert.Equal(t, user.Login, "login")
	assert.Equal(t, user.ID, 1)

	err := r.SaveUser(models.User{ID: 1, Login: "login_new"})
	assert.Nil(t, err)

	user = r.GetUserByLogin("login")
	assert.Nil(t, user)

	user = r.GetUserByLogin("login_new")
	assert.Equal(t, user.Login, "login_new")
	assert.Equal(t, user.ID, 1)
}
