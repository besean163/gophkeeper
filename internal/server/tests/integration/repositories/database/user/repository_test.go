package user

import (
	"testing"

	"github.com/besean163/gophkeeper/internal/server/models"
	"github.com/besean163/gophkeeper/internal/server/repositories/database/user"
	"github.com/stretchr/testify/assert"
)

func loadFixtureUsers(t *testing.T) {
	t.Helper()
	users := []models.User{
		{ID: 1, UUID: "00000000-0000-0000-0000-000000000001", Login: "login_1", Password: "password_1", CreatedAt: 1},
	}
	for _, user := range users {
		err := db.Create(&user).Error
		if err != nil {
			t.Fatalf("failed to load fixture: %s", err)
		}
	}
}

func cleanUpFixtureUsers(t *testing.T) {
	t.Helper()
	err := db.Exec("DELETE FROM users").Error
	if err != nil {
		t.Fatalf("failed to clean up fixture: %s", err)
	}
}

func TestGetUser(t *testing.T) {
	r := user.NewRepository(db)
	user, _ := r.GetUser(1)
	assert.Nil(t, user)

	loadFixtureUsers(t)
	defer cleanUpFixtureUsers(t)

	user, err := r.GetUser(1)

	assert.Nil(t, err)
	assert.NotNil(t, user)
}

func TestCreateUser(t *testing.T) {
	r := user.NewRepository(db)
	loadFixtureUsers(t)
	defer cleanUpFixtureUsers(t)

	user, _ := r.GetUser(2)
	assert.Nil(t, user)

	r.SaveUser(&models.User{ID: 2, UUID: "00000000-0000-0000-0000-000000000002", Login: "login_2", Password: "password_2", CreatedAt: 1})
	user, err := r.GetUser(2)

	assert.Nil(t, err)
	assert.NotNil(t, user)
}

func TestUpdateUser(t *testing.T) {
	r := user.NewRepository(db)
	loadFixtureUsers(t)
	defer cleanUpFixtureUsers(t)

	updatedUser := &models.User{ID: 1, UUID: "00000000-0000-0000-0000-000000000001", Login: "login_1", Password: "new_password", CreatedAt: 1}
	r.SaveUser(updatedUser)
	user, err := r.GetUser(1)

	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, user, updatedUser)
}

func TestCreateDuplicateUser(t *testing.T) {
	r := user.NewRepository(db)
	loadFixtureUsers(t)
	defer cleanUpFixtureUsers(t)

	err := r.SaveUser(&models.User{ID: 2, UUID: "00000000-0000-0000-0000-000000000001", Login: "login_1", Password: "password_1", CreatedAt: 1})
	assert.NotNil(t, err)
}

func TestUserByLogin(t *testing.T) {
	r := user.NewRepository(db)
	loadFixtureUsers(t)
	defer cleanUpFixtureUsers(t)

	var user *models.User
	user = r.GetUserByLogin("login_2")
	assert.Nil(t, user)

	user = r.GetUserByLogin("login_1")
	assert.NotNil(t, user)
}
