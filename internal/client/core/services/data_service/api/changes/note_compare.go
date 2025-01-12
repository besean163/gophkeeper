package changes

import (
	models "github.com/besean163/gophkeeper/internal/models/client"
	servermodels "github.com/besean163/gophkeeper/internal/models/server"
)

type NoteCompare struct {
	Items        []models.Note
	CompareItems []servermodels.Note
}
