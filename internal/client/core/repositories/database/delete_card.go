package database

import "github.com/besean163/gophkeeper/internal/client/core/models"

// DeleteCard удаление карты
func (r Repository) DeleteCard(uuid string) error {
	result := r.DB.Where("uuid = ?", uuid).Delete(&models.Card{})
	err := result.Error
	if err != nil {
		return err
	}
	return nil
}
