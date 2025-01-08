package database

import "github.com/besean163/gophkeeper/internal/client/core/models"

func (s Service) CreateCard(user models.User, item models.Card) error {
	if item.UUID == "" {
		item.UUID = s.uuidController.GetUUID()
	}

	item.UserID = user.ID
	item.CreatedAt = s.timeController.Now()
	item.UpdatedAt = s.timeController.Now()

	return s.repository.SaveCard(item)
}
