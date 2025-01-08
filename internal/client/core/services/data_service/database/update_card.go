package database

import "github.com/besean163/gophkeeper/internal/client/core/models"

func (s Service) UpdateCard(user models.User, item models.Card) error {
	if item.UUID == "" {
		item.UUID = s.uuidController.GetUUID()
	}

	item.UserID = user.ID
	item.UpdatedAt = s.timeController.Now()

	return s.repository.SaveCard(item)
}
