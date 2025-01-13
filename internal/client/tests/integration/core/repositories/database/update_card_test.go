package database

import (
	"testing"

	models "github.com/besean163/gophkeeper/internal/models/client"
	utilmock "github.com/besean163/gophkeeper/internal/tests/mocks/utils"
	"github.com/golang/mock/gomock"

	"github.com/besean163/gophkeeper/internal/client/core/repositories/database"
	"github.com/stretchr/testify/assert"
)

func TestUpdateCard(t *testing.T) {
	user := &models.User{ID: 1}
	item_1 := &models.Card{UUID: "00000000-0000-0000-0000-000000000001", UserID: 1, Name: "name_1", Number: 1111, Exp: "11|11", CVV: 111}

	loadFixtureCards(t, []*models.Card{
		item_1,
	})
	defer cleanUpFixtureCards(t)

	ctrl := gomock.NewController(t)
	uuidController := utilmock.NewMockUUIDController(ctrl)
	r := database.NewRepository(db, uuidController)
	var err error

	items, err := r.GetCards(*user)
	assert.Nil(t, err)
	assert.Equal(t, []models.Card{
		*item_1,
	}, items)

	item_change := &models.Card{UUID: "00000000-0000-0000-0000-000000000001", UserID: 1, Name: "new_name", Number: 9999, Exp: "99|99", CVV: 999}
	err = r.SaveCard(*item_change)
	assert.Nil(t, err)

	items, err = r.GetCards(*user)
	assert.Nil(t, err)
	assert.Equal(t, []models.Card{
		*item_change,
	}, items)
}
