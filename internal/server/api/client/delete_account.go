package client

import (
	"errors"
	"net/http"

	"github.com/besean163/gophkeeper/internal/server/api/entities"
)

// DeleteAccount запрос на удаление аккаунта.
// Параметры:
//   - input: структура запроса.
func (c Client) DeleteAccount(input entities.AccountDeleteInput) error {
	response, err := c.Delete(c.Host+"/api/account", input)

	if err != nil {
		return errors.New("request error")
	}

	if response.StatusCode() != http.StatusOK {
		return errors.New("server error")
	}

	return nil
}
