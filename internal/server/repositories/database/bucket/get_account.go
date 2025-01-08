package bucket

import "github.com/besean163/gophkeeper/internal/server/models"

// GetAccount получение аккаунта
func (r Repository) GetAccount(uuid string) (*models.Account, error) {
	var account *models.Account
	r.db.Where("uuid = ?", uuid).Find(&account)
	if account.ID == 0 {
		return nil, nil
	}
	return account, nil
}
