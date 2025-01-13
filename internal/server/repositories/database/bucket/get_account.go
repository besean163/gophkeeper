package bucket

import models "github.com/besean163/gophkeeper/internal/models/server"

// GetAccount получение аккаунта
func (r Repository) GetAccount(uuid string) (*models.Account, error) {
	var account *models.Account
	r.db.Where("uuid = ?", uuid).Find(&account)
	if account.UUID == "" {
		return nil, nil
	}
	return account, nil
}
