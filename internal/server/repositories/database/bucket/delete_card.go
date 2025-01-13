package bucket

import models "github.com/besean163/gophkeeper/internal/models/server"

// DeleteCard удаление карты
func (r Repository) DeleteCard(uuid string) error {
	result := r.db.Where("uuid = ?", uuid).Delete(&models.Card{})

	err := result.Error
	if err != nil {
		return err
	}
	return nil
}
