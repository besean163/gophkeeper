package client

import (
	"encoding/json"
	"errors"

	"github.com/besean163/gophkeeper/internal/server/api/entities/output"
)

// GetNotes запрос на получение списка заметок.
func (c Client) GetNotes() (*output.GetNotes, error) {
	var output *output.GetNotes
	response, err := c.Get(c.Host + "/api/notes")

	if err != nil {
		return nil, errors.New("request error")
	}

	err = json.Unmarshal(response.Body(), &output)
	if err != nil {
		return nil, errors.New("read answer error")
	}

	return output, nil
}
