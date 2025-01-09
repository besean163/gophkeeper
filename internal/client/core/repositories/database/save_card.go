package database

import (
	models "github.com/besean163/gophkeeper/internal/models/client"
)

// SaveCard сохранение карты
func (r Repository) SaveCard(item models.Card) error {
	if item.UUID == "" {
		item.UUID = r.UUIDController.GetUUID()
		return r.createItem(&item)
	}

	return r.updateItem(&item)
}
