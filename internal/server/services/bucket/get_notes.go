package bucket

import "github.com/besean163/gophkeeper/internal/server/models"

// GetNotes получение заметок
func (s Service) GetNotes(user models.User) ([]*models.Note, error) {
	items, err := s.repository.GetNotes(user)
	if err != nil {
		return nil, err
	}

	return items, nil
}
