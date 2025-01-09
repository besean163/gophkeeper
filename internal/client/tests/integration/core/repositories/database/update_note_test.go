package database

import (
	"testing"

	models "github.com/besean163/gophkeeper/internal/models/client"
	utilmock "github.com/besean163/gophkeeper/internal/tests/mocks/utils"
	"github.com/golang/mock/gomock"

	"github.com/besean163/gophkeeper/internal/client/core/repositories/database"
	"github.com/stretchr/testify/assert"
)

func TestUpdateNote(t *testing.T) {
	user := &models.User{ID: 1}
	item_1 := &models.Note{UUID: "00000000-0000-0000-0000-000000000001", UserID: 1, Name: "name_1", Content: "text_1"}

	loadFixtureNotes(t, []*models.Note{
		item_1,
	})
	defer cleanUpFixtureNotes(t)

	ctrl := gomock.NewController(t)
	uuidController := utilmock.NewMockUUIDController(ctrl)
	r := database.NewRepository(db, uuidController)
	var err error

	items, err := r.GetNotes(*user)
	assert.Nil(t, err)
	assert.Equal(t, []models.Note{
		*item_1,
	}, items)

	item_change := &models.Note{UUID: "00000000-0000-0000-0000-000000000001", UserID: 1, Name: "new_name", Content: "new_text"}
	err = r.SaveNote(*item_change)
	assert.Nil(t, err)

	items, err = r.GetNotes(*user)
	assert.Nil(t, err)
	assert.Equal(t, []models.Note{
		*item_change,
	}, items)
}
