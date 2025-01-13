package database

import (
	models "github.com/besean163/gophkeeper/internal/models/client"
)

func (s Service) CreateNote(user models.User, item models.Note) error {
	item.UserID = user.ID
	item.CreatedAt = s.timeController.Now()
	item.UpdatedAt = s.timeController.Now()

	return s.repository.SaveNote(item)
}
