package database

import (
	models "github.com/besean163/gophkeeper/internal/models/client"
)

// GetNotes получение списка заметок
func (r Repository) GetNotes(user models.User) ([]models.Note, error) {
	items := []models.Note{}
	result := r.DB.Where("user_id = ?", user.ID).Find(&items)
	err := result.Error
	if err != nil {
		return nil, err
	}
	return items, nil
}
