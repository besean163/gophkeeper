package client

import (
	"errors"
	"net/http"

	"github.com/besean163/gophkeeper/internal/server/api/entities/input"
)

// UpdateAccount запрос на обновление аккаунта.
// Параметры:
//   - input: структура запроса.
func (c Client) UpdateAccount(input input.AccountUpdate) error {
	response, err := c.Put(c.Host+"/api/account", input)

	if err != nil {
		return errors.New("request error")
	}

	if response.StatusCode() != http.StatusOK {
		return errors.New("server error")
	}

	return nil
}
