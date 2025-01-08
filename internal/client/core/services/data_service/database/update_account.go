package database

import "github.com/besean163/gophkeeper/internal/client/core/models"

func (s Service) UpdateAccount(user models.User, item models.Account) error {
	if item.UUID == "" {
		item.UUID = s.uuidController.GetUUID()
	}

	item.UserID = user.ID
	item.UpdatedAt = s.timeController.Now()

	return s.repository.SaveAccount(item)
}
