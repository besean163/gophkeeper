package entities

import "github.com/besean163/gophkeeper/internal/server/models"

// GetAccountsOutput структура для получения аккаунтов.
type GetAccountsOutput struct {
	Accounts []*models.Account `json:"accounts"`
}
