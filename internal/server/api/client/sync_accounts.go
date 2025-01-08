package client

import (
	"net/http"

	"github.com/besean163/gophkeeper/internal/server/api/entities"
)

// SyncAccounts запрос на синхронизацию аккаунтов.
// Параметры:
//   - input: структура запроса.
func (c Client) SyncAccounts(input entities.AccountsSyncInput) error {
	response, err := c.Post(c.Host+"/api/accounts/sync", input)

	if err != nil {
		return ErrorRequestError
	}

	if response.StatusCode() != http.StatusOK {
		return ErrorServerError
	}

	return nil
}
