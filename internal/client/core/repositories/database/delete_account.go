package database

import (
	models "github.com/besean163/gophkeeper/internal/models/client"
)

// DeleteAccount удаление аккаунта
func (r Repository) DeleteAccount(uuid string) error {
	result := r.DB.Where("uuid = ?", uuid).Delete(&models.Account{})
	err := result.Error
	if err != nil {
		return err
	}
	return nil
}
