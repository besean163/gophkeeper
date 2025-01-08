package client

import (
	"net/http"

	"github.com/besean163/gophkeeper/internal/server/api/entities"
)

// SyncNotes запрос на синхронизацию заметок.
// Параметры:
//   - input: структура запроса.
func (c Client) SyncNotes(input entities.NotesSyncInput) error {
	response, err := c.Post(c.Host+"/api/notes/sync", input)

	if err != nil {
		return ErrorRequestError
	}

	if response.StatusCode() != http.StatusOK {
		return ErrorServerError
	}

	return nil
}
