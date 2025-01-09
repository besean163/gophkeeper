package bucket

import (
	models "github.com/besean163/gophkeeper/internal/models/server"
)

// CreateNote создание заметки
func (s Service) CreateNote(user models.User, item *models.Note) error {
	if item.UUID == "" {
		item.UUID = s.uuidController.GetUUID()
	}

	if item.UserID == 0 {
		item.UserID = user.ID
	}

	item.CreatedAt = s.timeController.Now()

	return s.repository.SaveNote(item)
}
