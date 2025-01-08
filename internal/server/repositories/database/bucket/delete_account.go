package bucket

import "github.com/besean163/gophkeeper/internal/server/models"

// DeleteAccount удаление аккаунта
func (r Repository) DeleteAccount(uuid string) error {
	result := r.db.Where("uuid = ?", uuid).Delete(&models.Account{})

	err := result.Error
	if err != nil {
		return err
	}
	return nil
}
