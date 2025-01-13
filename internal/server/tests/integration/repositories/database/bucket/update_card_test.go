package bucket

import (
	"testing"

	models "github.com/besean163/gophkeeper/internal/models/server"
	"github.com/besean163/gophkeeper/internal/server/repositories/database/bucket"
	utilmock "github.com/besean163/gophkeeper/internal/tests/mocks/utils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUpdateCard(t *testing.T) {
	user := models.User{ID: 1}
	item_1 := &models.Card{UUID: "00000000-0000-0000-0000-000000000001", UserID: 1, Name: "name_1", Number: 1111, Exp: "11|11", CVV: 111, CreatedAt: 1, UpdatedAt: 1}
	item_2 := &models.Card{UUID: "00000000-0000-0000-0000-000000000002", UserID: 2, Name: "name_2", Number: 2222, Exp: "22|22", CVV: 222, CreatedAt: 1, UpdatedAt: 1}

	loadFixtureCards(t, []*models.Card{
		item_1,
		item_2,
	})
	defer cleanUpFixtureCards(t)

	ctrl := gomock.NewController(t)
	uuidController := utilmock.NewMockUUIDController(ctrl)
	r := bucket.NewRepository(db, uuidController)

	var err error
	item := models.Card{UserID: user.ID, UUID: "00000000-0000-0000-0000-000000000001", Name: "name_1_new", Number: 9999, Exp: "99|99", CVV: 999, CreatedAt: 1, UpdatedAt: 2}

	err = r.SaveCard(&item)
	assert.Nil(t, err)

	var updateItem *models.Card
	updateItem, err = r.GetCard("00000000-0000-0000-0000-000000000001")
	assert.Nil(t, err)
	assert.Equal(t, &models.Card{UserID: user.ID, UUID: "00000000-0000-0000-0000-000000000001", Name: "name_1_new", Number: 9999, Exp: "99|99", CVV: 999, CreatedAt: 1, UpdatedAt: 2}, updateItem)
}
