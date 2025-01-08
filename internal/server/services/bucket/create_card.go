package bucket

import (
	"github.com/besean163/gophkeeper/internal/server/models"
)

// CreateCard создание карты
func (s Service) CreateCard(user models.User, item *models.Card) error {
	if item.UUID == "" {
		item.UUID = s.uuidController.GetUUID()
	}

	if item.UserID == 0 {
		item.UserID = user.ID
	}

	item.CreatedAt = s.timeController.Now()

	return s.repository.SaveCard(item)
}
