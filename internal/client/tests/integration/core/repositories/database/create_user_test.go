package database

import (
	"testing"

	models "github.com/besean163/gophkeeper/internal/models/client"
	utilmock "github.com/besean163/gophkeeper/internal/tests/mocks/utils"
	"github.com/golang/mock/gomock"

	"github.com/besean163/gophkeeper/internal/client/core/repositories/database"
	"github.com/stretchr/testify/assert"
)

func TestCreateUser(t *testing.T) {
	user_1 := &models.User{ID: 1, Login: "login"}
	loadFixtureUsers(t, []*models.User{
		user_1,
	})
	defer cleanUpFixtureUser(t)

	ctrl := gomock.NewController(t)
	uuidController := utilmock.NewMockUUIDController(ctrl)

	r := database.NewRepository(db, uuidController)
	var user *models.User

	user = r.GetUserByLogin("login_2")
	assert.Nil(t, user)

	user_2 := &models.User{ID: 2, Login: "login_2"}
	err := r.SaveUser(*user_2)
	assert.Nil(t, err)

	user = r.GetUserByLogin("login_2")
	assert.NotNil(t, user)
}

func TestCreateDuplicateUser(t *testing.T) {
	user_1 := &models.User{ID: 1, Login: "login"}
	loadFixtureUsers(t, []*models.User{
		user_1,
	})
	defer cleanUpFixtureUser(t)

	ctrl := gomock.NewController(t)
	uuidController := utilmock.NewMockUUIDController(ctrl)
	r := database.NewRepository(db, uuidController)

	user_2 := &models.User{ID: 2, Login: "login"}
	err := r.SaveUser(*user_2)
	assert.NotNil(t, err)
}
