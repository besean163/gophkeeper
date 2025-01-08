package database

import "github.com/besean163/gophkeeper/internal/client/core/models"

// DeleteAccount удаление аккаунта
func (r Repository) DeleteAccount(uuid string) error {
	result := r.DB.Where("uuid = ?", uuid).Delete(&models.Account{})
	err := result.Error
	if err != nil {
		return err
	}
	return nil
}
