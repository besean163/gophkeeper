package bucket

import (
	"github.com/besean163/gophkeeper/internal/server/models"
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
