package bucket

import (
	"testing"

	"github.com/besean163/gophkeeper/internal/server/models"
	"github.com/besean163/gophkeeper/internal/server/repositories/database/bucket"
	"github.com/stretchr/testify/assert"
)

func TestDeleteNote(t *testing.T) {
	item_1 := &models.Note{ID: 1, UUID: "00000000-0000-0000-0000-000000000001", UserID: 1, Name: "name_1", Content: "text_1", CreatedAt: 1, UpdatedAt: 1}
	item_2 := &models.Note{ID: 2, UUID: "00000000-0000-0000-0000-000000000002", UserID: 2, Name: "name_2", Content: "text_2", CreatedAt: 1, UpdatedAt: 1}

	loadFixtureNotes(t, []*models.Note{
		item_1,
		item_2,
	})
	defer cleanUpFixtureNotes(t)

	r := bucket.NewRepository(db)

	var err error
	var item *models.Note

	item, err = r.GetNote("00000000-0000-0000-0000-000000000001")
	assert.Nil(t, err)
	assert.Equal(t, item_1, item)

	err = r.DeleteNote(item.UUID)
	assert.Nil(t, err)

	item, err = r.GetNote("00000000-0000-0000-0000-000000000001")
	assert.Nil(t, err)
	assert.Nil(t, item)
}
