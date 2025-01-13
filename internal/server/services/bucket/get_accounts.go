package bucket

import models "github.com/besean163/gophkeeper/internal/models/server"

// GetAccounts получение аккаунтов
func (s Service) GetAccounts(user models.User) ([]*models.Account, error) {
	items, err := s.repository.GetAccounts(user)
	if err != nil {
		return nil, err
	}

	return items, nil
}
