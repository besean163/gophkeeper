package bucket

import (
	"testing"

	"github.com/besean163/gophkeeper/internal/server/models"
	"github.com/besean163/gophkeeper/internal/server/repositories/database/bucket"
	"github.com/stretchr/testify/assert"
)

func TestUpdateCard(t *testing.T) {
	user := models.User{ID: 1}
	item_1 := &models.Card{ID: 1, UUID: "00000000-0000-0000-0000-000000000001", UserID: 1, Name: "name_1", Number: 1111, Exp: "11|11", CVV: 111, CreatedAt: 1, UpdatedAt: 1}
	item_2 := &models.Card{ID: 2, UUID: "00000000-0000-0000-0000-000000000002", UserID: 2, Name: "name_2", Number: 2222, Exp: "22|22", CVV: 222, CreatedAt: 1, UpdatedAt: 1}

	loadFixtureCards(t, []*models.Card{
		item_1,
		item_2,
	})
	defer cleanUpFixtureCards(t)

	r := bucket.NewRepository(db)

	var err error
	item := models.Card{ID: 1, UserID: user.ID, UUID: "00000000-0000-0000-0000-000000000001", Name: "name_1_new", Number: 9999, Exp: "99|99", CVV: 999, CreatedAt: 1, UpdatedAt: 2}

	err = r.SaveCard(&item)
	assert.Nil(t, err)

	var updateItem *models.Card
	updateItem, err = r.GetCard("00000000-0000-0000-0000-000000000001")
	assert.Nil(t, err)
	assert.Equal(t, &models.Card{ID: 1, UserID: user.ID, UUID: "00000000-0000-0000-0000-000000000001", Name: "name_1_new", Number: 9999, Exp: "99|99", CVV: 999, CreatedAt: 1, UpdatedAt: 2}, updateItem)
}
