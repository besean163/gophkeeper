package database

import (
	"testing"

	"github.com/besean163/gophkeeper/internal/client/core/models"
	"github.com/besean163/gophkeeper/internal/client/core/repositories/database"
	"github.com/stretchr/testify/assert"
)

func loadFixtureUsers(t *testing.T, users []*models.User) {
	t.Helper()
	for _, user := range users {
		err := db.Create(&user).Error
		if err != nil {
			t.Fatalf("failed to load fixture: %s", err)
		}
	}
}

func cleanUpFixtureUser(t *testing.T) {
	t.Helper()
	err := db.Exec("DELETE FROM users").Error
	if err != nil {
		t.Fatalf("failed to clean up fixture: %s", err)
	}
}

func loadFixtureAccounts(t *testing.T, accounts []*models.Account) {
	t.Helper()
	for _, account := range accounts {
		err := db.Create(&account).Error
		if err != nil {
			t.Fatalf("failed to load fixture: %s", err)
		}
	}
}

func cleanUpFixtureAccounts(t *testing.T) {
	t.Helper()
	err := db.Exec("DELETE FROM accounts").Error
	if err != nil {
		t.Fatalf("failed to clean up fixture: %s", err)
	}
}

func TestGetUserByLogin(t *testing.T) {
	user_1 := &models.User{ID: 1, Login: "login"}

	loadFixtureUsers(t, []*models.User{
		user_1,
	})
	defer cleanUpFixtureUser(t)

	r := database.NewRepository(db)
	var user *models.User
	user = r.GetUserByLogin("wrong_login")
	assert.Nil(t, user)

	user = r.GetUserByLogin("login")
	assert.NotNil(t, user)
}

func TestCreateUser(t *testing.T) {
	user_1 := &models.User{ID: 1, Login: "login"}
	loadFixtureUsers(t, []*models.User{
		user_1,
	})
	defer cleanUpFixtureUser(t)

	r := database.NewRepository(db)
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

	r := database.NewRepository(db)

	user_2 := &models.User{ID: 2, Login: "login"}
	err := r.SaveUser(*user_2)
	assert.NotNil(t, err)
}

func TestUpdateUser(t *testing.T) {
	user_1 := &models.User{ID: 1, Login: "login"}
	loadFixtureUsers(t, []*models.User{
		user_1,
	})
	defer cleanUpFixtureUser(t)

	r := database.NewRepository(db)
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

func TestGetAccounts(t *testing.T) {
	user := &models.User{ID: 1}
	account_1 := &models.Account{UUID: "00000000-0000-0000-0000-000000000001", UserID: 1, Name: "name_1", Login: "login_1", Password: "password_1"}
	account_2 := &models.Account{UUID: "00000000-0000-0000-0000-000000000002", UserID: 2, Name: "name_2", Login: "login_2", Password: "password_2"}
	account_3 := &models.Account{UUID: "00000000-0000-0000-0000-000000000003", UserID: 1, Name: "name_3", Login: "login_3", Password: "password_3"}

	loadFixtureAccounts(t, []*models.Account{
		account_1,
		account_2,
		account_3,
	})
	defer cleanUpFixtureAccounts(t)

	r := database.NewRepository(db)
	accounts, err := r.GetAccounts(*user)
	assert.Nil(t, err)

	assert.Equal(t, []models.Account{
		*account_1,
		*account_3,
	}, accounts)
}

func TestCreateAccount(t *testing.T) {
	user := &models.User{ID: 1}
	account_1 := &models.Account{ID: 1, UUID: "00000000-0000-0000-0000-000000000001", UserID: 1, Name: "name_1", Login: "login_1", Password: "password_1"}
	account_2 := &models.Account{ID: 2, UUID: "00000000-0000-0000-0000-000000000002", UserID: 2, Name: "name_2", Login: "login_2", Password: "password_2"}
	account_3 := &models.Account{ID: 3, UUID: "00000000-0000-0000-0000-000000000003", UserID: 1, Name: "name_3", Login: "login_3", Password: "password_3"}

	loadFixtureAccounts(t, []*models.Account{
		account_1,
		account_2,
		account_3,
	})
	defer cleanUpFixtureAccounts(t)

	r := database.NewRepository(db)
	var err error

	accounts, err := r.GetAccounts(*user)
	assert.Nil(t, err)
	assert.Equal(t, []models.Account{
		*account_1,
		*account_3,
	}, accounts)

	account_4 := &models.Account{ID: 4, UUID: "00000000-0000-0000-0000-000000000004", UserID: 1, Name: "name_4", Login: "login_4", Password: "password_4"}
	err = r.SaveAccount(*account_4)
	assert.Nil(t, err)

	accounts, err = r.GetAccounts(*user)
	assert.Nil(t, err)
	assert.Equal(t, []models.Account{
		*account_1,
		*account_3,
		*account_4,
	}, accounts)
}

func TestUpdateAccount(t *testing.T) {
	user := &models.User{ID: 1}
	account_1 := &models.Account{ID: 1, UUID: "00000000-0000-0000-0000-000000000001", UserID: 1, Name: "name_1", Login: "login_1", Password: "password_1"}

	loadFixtureAccounts(t, []*models.Account{
		account_1,
	})
	defer cleanUpFixtureAccounts(t)

	r := database.NewRepository(db)
	var err error

	accounts, err := r.GetAccounts(*user)
	assert.Nil(t, err)
	assert.Equal(t, []models.Account{
		*account_1,
	}, accounts)

	account_change := &models.Account{ID: 1, UUID: "00000000-0000-0000-0000-000000000001", UserID: 1, Name: "new_name", Login: "new_login", Password: "password_1"}
	err = r.SaveAccount(*account_change)
	assert.Nil(t, err)

	accounts, err = r.GetAccounts(*user)
	assert.Nil(t, err)
	assert.Equal(t, []models.Account{
		*account_change,
	}, accounts)
}

func TestDeleteAccount(t *testing.T) {
	user := &models.User{ID: 1}
	account_1 := &models.Account{ID: 1, UUID: "00000000-0000-0000-0000-000000000001", UserID: 1, Name: "name_1", Login: "login_1", Password: "password_1"}
	account_2 := &models.Account{ID: 2, UUID: "00000000-0000-0000-0000-000000000002", UserID: 2, Name: "name_2", Login: "login_2", Password: "password_2"}
	account_3 := &models.Account{ID: 3, UUID: "00000000-0000-0000-0000-000000000003", UserID: 1, Name: "name_3", Login: "login_3", Password: "password_3"}

	loadFixtureAccounts(t, []*models.Account{
		account_1,
		account_2,
		account_3,
	})
	defer cleanUpFixtureAccounts(t)

	r := database.NewRepository(db)
	var err error

	accounts, err := r.GetAccounts(*user)
	assert.Nil(t, err)
	assert.Equal(t, []models.Account{
		*account_1,
		*account_3,
	}, accounts)

	err = r.DeleteAccount(account_1.UUID)
	assert.Nil(t, err)

	accounts, err = r.GetAccounts(*user)
	assert.Nil(t, err)
	assert.Equal(t, []models.Account{
		*account_3,
	}, accounts)
}
