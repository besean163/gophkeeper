package client

import (
	"errors"
	"net/http"

	"github.com/besean163/gophkeeper/internal/server/api/entities/input"
)

// UpdateCard запрос на обновление карты.
// Параметры:
//   - input: структура запроса.
func (c Client) UpdateCard(input input.CardUpdate) error {
	response, err := c.Put(c.Host+"/api/card", input)

	if err != nil {
		return errors.New("request error")
	}

	if response.StatusCode() != http.StatusOK {
		return errors.New("server error")
	}

	return nil
}
