package entities

import "github.com/besean163/gophkeeper/internal/server/models"

// GetCardsOutput структура для получения карт.
type GetCardsOutput struct {
	Cards []*models.Card `json:"cards"`
}
