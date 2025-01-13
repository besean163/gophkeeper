package bucket

import models "github.com/besean163/gophkeeper/internal/models/server"

// GetCards получение карт
func (s Service) GetCards(user models.User) ([]*models.Card, error) {
	items, err := s.repository.GetCards(user)
	if err != nil {
		return nil, err
	}

	return items, nil
}
