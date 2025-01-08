package entities

import "github.com/besean163/gophkeeper/internal/server/models"

// GetNotesOutput структура для получения заметок.
type GetNotesOutput struct {
	Notes []*models.Note `json:"notes"`
}
