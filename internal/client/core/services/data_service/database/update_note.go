package database

import (
	models "github.com/besean163/gophkeeper/internal/models/client"
)

func (s Service) UpdateNote(user models.User, item models.Note) error {
	item.UserID = user.ID
	item.UpdatedAt = s.timeController.Now()

	return s.repository.SaveNote(item)
}
