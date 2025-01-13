package changes

import (
	models "github.com/besean163/gophkeeper/internal/models/server"
)

type NoteChanges struct {
	Created []*models.Note
	Updated []*models.Note
	Deleted []string
}

func NewNoteChanges() NoteChanges {
	return NoteChanges{
		Created: make([]*models.Note, 0),
		Updated: make([]*models.Note, 0),
		Deleted: make([]string, 0),
	}
}
