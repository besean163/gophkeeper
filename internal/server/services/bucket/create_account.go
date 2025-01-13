package bucket

import (
	models "github.com/besean163/gophkeeper/internal/models/server"
)

// CreateAccount создание аккаунта
func (s Service) CreateAccount(user models.User, item *models.Account) error {
	if item.UUID == "" {
		item.UUID = s.uuidController.GetUUID()
	}

	if item.UserID == 0 {
		item.UserID = user.ID
	}

	item.CreatedAt = s.timeController.Now()
	return s.repository.SaveAccount(item)
}
