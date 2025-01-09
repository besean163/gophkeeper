package output

import (
	models "github.com/besean163/gophkeeper/internal/models/server"
)

// GetNotes структура для получения заметок.
type GetNotes struct {
	Notes []*models.Note `json:"notes"`
}
