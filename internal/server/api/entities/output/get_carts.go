package output

import (
	models "github.com/besean163/gophkeeper/internal/models/server"
)

// GetCards структура для получения карт.
type GetCards struct {
	Cards []*models.Card `json:"cards"`
}
