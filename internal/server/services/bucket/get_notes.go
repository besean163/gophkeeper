package bucket

import models "github.com/besean163/gophkeeper/internal/models/server"

// GetNotes получение заметок
func (s Service) GetNotes(user models.User) ([]*models.Note, error) {
	items, err := s.repository.GetNotes(user)
	if err != nil {
		return nil, err
	}

	return items, nil
}
