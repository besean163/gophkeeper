package database

import (
	models "github.com/besean163/gophkeeper/internal/models/client"
)

// GetAccounts получение списка аккаунтов
func (r Repository) GetAccounts(user models.User) ([]models.Account, error) {
	items := []models.Account{}
	result := r.DB.Where("user_id = ?", user.ID).Find(&items)
	err := result.Error
	if err != nil {
		return nil, err
	}
	return items, nil
}
