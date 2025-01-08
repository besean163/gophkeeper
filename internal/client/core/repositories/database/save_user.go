package database

import "github.com/besean163/gophkeeper/internal/client/core/models"

// SaveUser сохранение пользователя
func (r Repository) SaveUser(user models.User) error {
	if user.ID == 0 {
		return r.insertItem(&user)
	}

	return r.updateItem(&user)
}
