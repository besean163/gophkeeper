package database

import (
	models "github.com/besean163/gophkeeper/internal/models/client"
)

// GetUserByLogin получение пользователя по логину
func (r Repository) GetUserByLogin(login string) *models.User {
	user := new(models.User)
	r.DB.Where("login = ?", login).First(user)

	if user.ID == 0 {
		return nil
	}

	return user
}
