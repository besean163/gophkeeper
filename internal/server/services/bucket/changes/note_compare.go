package changes

import (
	clientmodels "github.com/besean163/gophkeeper/internal/models/client"
	models "github.com/besean163/gophkeeper/internal/models/server"
)

type NoteCompare struct {
	Items        []*models.Note
	CompareItems []clientmodels.Note
}
