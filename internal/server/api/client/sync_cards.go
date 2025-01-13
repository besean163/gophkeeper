package client

import (
	"net/http"

	"github.com/besean163/gophkeeper/internal/server/api/entities/input"
)

// SyncCards запрос на синхронизацию карт.
// Параметры:
//   - input: структура запроса.
func (c Client) SyncCards(input input.CardsSync) error {
	response, err := c.Post(c.Host+"/api/cards/sync", input)

	if err != nil {
		return ErrorRequestError
	}

	if response.StatusCode() != http.StatusOK {
		return ErrorServerError
	}

	return nil
}
