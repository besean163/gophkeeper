package client

import (
	"errors"
	"net/http"

	"github.com/besean163/gophkeeper/internal/server/api/entities/input"
)

// CreateCard запрос на создание карты.
// Параметры:
//   - input: структура запроса.
func (c Client) CreateCard(input input.CardCreate) error {
	response, err := c.Post(c.Host+"/api/card", input)

	if err != nil {
		return errors.New("request error")
	}

	if response.StatusCode() != http.StatusOK {
		return errors.New("server error")
	}

	return nil
}
