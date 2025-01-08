package database

import "github.com/besean163/gophkeeper/internal/client/core/models"

// SaveCard сохранение карты
func (r Repository) SaveCard(item models.Card) error {
	if item.ID == 0 {
		return r.insertItem(&item)
	}

	return r.updateItem(&item)
}
