package client

import (
	"errors"
	"net/http"

	"github.com/besean163/gophkeeper/internal/server/api/entities/input"
)

// DeleteCard запрос на удаление карты.
// Параметры:
//   - input: структура запроса.
func (c Client) DeleteCard(input input.CardDelete) error {
	response, err := c.Delete(c.Host+"/api/card", input)

	if err != nil {
		return errors.New("request error")
	}

	if response.StatusCode() != http.StatusOK {
		return errors.New("server error")
	}

	return nil
}
