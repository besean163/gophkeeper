package database

import "github.com/besean163/gophkeeper/internal/client/core/models"

func (s Service) CreateNote(user models.User, item models.Note) error {
	if item.UUID == "" {
		item.UUID = s.uuidController.GetUUID()
	}

	item.UserID = user.ID
	item.CreatedAt = s.timeController.Now()
	item.UpdatedAt = s.timeController.Now()

	return s.repository.SaveNote(item)
}
