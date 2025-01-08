package client

import (
	"errors"
	"net/http"

	"github.com/besean163/gophkeeper/internal/server/api/entities"
)

// CreateNote запрос на создание заметки.
// Параметры:
//   - input: структура запроса.
func (c Client) CreateNote(input entities.NoteCreateInput) error {
	response, err := c.Post(c.Host+"/api/note", input)

	if err != nil {
		return errors.New("request error")
	}

	if response.StatusCode() != http.StatusOK {
		return errors.New("server error")
	}

	return nil
}
