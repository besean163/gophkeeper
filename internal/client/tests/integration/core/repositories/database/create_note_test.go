package database

import (
	"testing"

	models "github.com/besean163/gophkeeper/internal/models/client"
	utilmock "github.com/besean163/gophkeeper/internal/tests/mocks/utils"
	"github.com/golang/mock/gomock"

	"github.com/besean163/gophkeeper/internal/client/core/repositories/database"
	"github.com/stretchr/testify/assert"
)

func TestCreateNote(t *testing.T) {
	user := &models.User{ID: 1}
	item_1 := &models.Note{UUID: "00000000-0000-0000-0000-000000000001", UserID: 1, Name: "name_1", Content: "text_1"}
	item_2 := &models.Note{UUID: "00000000-0000-0000-0000-000000000002", UserID: 2, Name: "name_2", Content: "text_2"}
	item_3 := &models.Note{UUID: "00000000-0000-0000-0000-000000000003", UserID: 1, Name: "name_3", Content: "text_3"}

	loadFixtureNotes(t, []*models.Note{
		item_1,
		item_2,
		item_3,
	})
	defer cleanUpFixtureNotes(t)

	ctrl := gomock.NewController(t)
	uuidController := utilmock.NewMockUUIDController(ctrl)
	uuidController.EXPECT().GetUUID().Return("00000000-0000-0000-0000-000000000004").Times(1)

	r := database.NewRepository(db, uuidController)
	var err error

	items, err := r.GetNotes(*user)
	assert.Nil(t, err)
	assert.Equal(t, []models.Note{
		*item_1,
		*item_3,
	}, items)

	item_4 := &models.Note{UserID: 1, Name: "name_4", Content: "text_4"}
	err = r.SaveNote(*item_4)
	assert.Nil(t, err)

	items, err = r.GetNotes(*user)
	assert.Nil(t, err)
	assert.Equal(t, []models.Note{
		*item_1,
		*item_3,
		{
			UUID:    "00000000-0000-0000-0000-000000000004",
			UserID:  1,
			Name:    "name_4",
			Content: "text_4",
		},
	}, items)
}
