package output

import (
	models "github.com/besean163/gophkeeper/internal/models/server"
)

// GetAccounts структура для получения аккаунтов.
type GetAccounts struct {
	Accounts []*models.Account `json:"accounts"`
}
