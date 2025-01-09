package database

import (
	models "github.com/besean163/gophkeeper/internal/models/client"
)

// GetCards получение списка карт
func (r Repository) GetCards(user models.User) ([]models.Card, error) {
	items := []models.Card{}
	result := r.DB.Where("user_id = ?", user.ID).Find(&items)
	err := result.Error
	if err != nil {
		return nil, err
	}
	return items, nil
}
