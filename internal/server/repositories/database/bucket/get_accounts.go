package bucket

import models "github.com/besean163/gophkeeper/internal/models/server"

// GetAccounts получение списка аккаунтов
func (r Repository) GetAccounts(user models.User) ([]*models.Account, error) {
	accounts := []*models.Account{}
	r.db.Where("user_id = ?", user.ID).Find(&accounts)
	return accounts, nil
}
