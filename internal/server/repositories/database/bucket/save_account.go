package bucket

import (
	models "github.com/besean163/gophkeeper/internal/models/server"
)

// SaveAccount сохранение аккаунта
func (r Repository) SaveAccount(item *models.Account) error {
	if item.UUID == "" {
		item.UUID = r.UUIDController.GetUUID()
		return r.createItem(&item)
	}

	return r.updateItem(&item)
}