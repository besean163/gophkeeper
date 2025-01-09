package client

import (
	"errors"
	"net/http"

	"github.com/besean163/gophkeeper/internal/server/api/entities/input"
)

// UpdateNote запрос на обновление заметки.
// Параметры:
//   - input: структура запроса.
func (c Client) UpdateNote(input input.NoteUpdate) error {
	response, err := c.Put(c.Host+"/api/note", input)

	if err != nil {
		return errors.New("request error")
	}

	if response.StatusCode() != http.StatusOK {
		return errors.New("server error")
	}

	return nil
}
