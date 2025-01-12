package bucket

import (
	"testing"

	models "github.com/besean163/gophkeeper/internal/models/server"
	"github.com/besean163/gophkeeper/internal/server/repositories/database/bucket"
	utilmock "github.com/besean163/gophkeeper/internal/tests/mocks/utils"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUpdateNote(t *testing.T) {
	user := models.User{ID: 1}
	item_1 := &models.Note{UUID: "00000000-0000-0000-0000-000000000001", UserID: 1, Name: "name_1", Content: "text_1", CreatedAt: 1, UpdatedAt: 1}
	item_2 := &models.Note{UUID: "00000000-0000-0000-0000-000000000002", UserID: 2, Name: "name_2", Content: "text_2", CreatedAt: 1, UpdatedAt: 1}

	loadFixtureNotes(t, []*models.Note{
		item_1,
		item_2,
	})
	defer cleanUpFixtureNotes(t)

	ctrl := gomock.NewController(t)
	uuidController := utilmock.NewMockUUIDController(ctrl)
	r := bucket.NewRepository(db, uuidController)

	var err error
	item := models.Note{UserID: user.ID, UUID: "00000000-0000-0000-0000-000000000001", Name: "name_1_new", Content: "text_1_new", CreatedAt: 1, UpdatedAt: 2}

	err = r.SaveNote(&item)
	assert.Nil(t, err)

	var updateItem *models.Note
	updateItem, err = r.GetNote("00000000-0000-0000-0000-000000000001")
	assert.Nil(t, err)
	assert.Equal(t, &models.Note{UserID: user.ID, UUID: "00000000-0000-0000-0000-000000000001", Name: "name_1_new", Content: "text_1_new", CreatedAt: 1, UpdatedAt: 2}, updateItem)
}