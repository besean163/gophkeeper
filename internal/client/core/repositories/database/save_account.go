package database

import "github.com/besean163/gophkeeper/internal/client/core/models"

// SaveAccount сохранение аккаунта
func (r Repository) SaveAccount(item models.Account) error {
	if item.ID == 0 {
		return r.insertItem(&item)
	}

	return r.updateItem(&item)
}
