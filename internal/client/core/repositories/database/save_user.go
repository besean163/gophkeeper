package database

import (
	models "github.com/besean163/gophkeeper/internal/models/client"
)

// SaveUser сохранение пользователя
func (r Repository) SaveUser(item models.User) error {
	if item.ID == 0 {
		return r.createItem(&item)
	}

	return r.updateItem(&item)
}
