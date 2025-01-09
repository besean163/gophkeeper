package client

import (
	"net/http"

	"github.com/besean163/gophkeeper/internal/server/api/entities/input"
)

// SyncNotes запрос на синхронизацию заметок.
// Параметры:
//   - input: структура запроса.
func (c Client) SyncNotes(input input.NotesSync) error {
	response, err := c.Post(c.Host+"/api/notes/sync", input)

	if err != nil {
		return ErrorRequestError
	}

	if response.StatusCode() != http.StatusOK {
		return ErrorServerError
	}

	return nil
}
