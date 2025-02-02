package client

import (
	"net/http"

	"github.com/besean163/gophkeeper/internal/server/api/entities/input"
)

// CreateAccount запрос на создание аккаунта.
// Параметры:
//   - input: структура запроса.
func (c Client) CreateAccount(input input.AccountCreate) error {
	response, err := c.Post(c.Host+"/api/account", input)

	if err != nil {
		return ErrorRequestError
	}

	if response.StatusCode() != http.StatusOK {
		return ErrorServerError
	}

	return nil
}
