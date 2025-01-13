package client

import (
	"errors"
	"net/http"

	"github.com/besean163/gophkeeper/internal/server/api/entities/input"
)

// DeleteNote запрос на удаление заметки.
// Параметры:
//   - input: структура запроса.
func (c Client) DeleteNote(input input.NoteDelete) error {
	response, err := c.Delete(c.Host+"/api/note", input)

	if err != nil {
		return errors.New("request error")
	}

	if response.StatusCode() != http.StatusOK {
		return errors.New("server error")
	}

	return nil
}
