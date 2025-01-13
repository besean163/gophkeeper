package client

import (
	"encoding/json"
	"errors"

	"github.com/besean163/gophkeeper/internal/server/api/entities/output"
)

// GetCards запрос на получение списка карт.
func (c Client) GetCards() (*output.GetCards, error) {
	var output *output.GetCards
	response, err := c.Get(c.Host + "/api/cards")

	if err != nil {
		return nil, errors.New("request error")
	}

	err = json.Unmarshal(response.Body(), &output)
	if err != nil {
		return nil, errors.New("read answer error")
	}

	return output, nil
}
