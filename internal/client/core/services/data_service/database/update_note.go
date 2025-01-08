package database

import "github.com/besean163/gophkeeper/internal/client/core/models"

func (s Service) UpdateNote(user models.User, item models.Note) error {
	if item.UUID == "" {
		item.UUID = s.uuidController.GetUUID()
	}

	item.UserID = user.ID
	item.UpdatedAt = s.timeController.Now()

	return s.repository.SaveNote(item)
}
