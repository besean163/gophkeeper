package bucket

import (
	models "github.com/besean163/gophkeeper/internal/models/server"
	apierrors "github.com/besean163/gophkeeper/internal/server/api/errors"
)

// UpdateAccount обновление аккаунта
func (s Service) UpdateAccount(user models.User, item *models.Account) error {
	exist, err := s.repository.GetAccount(item.UUID)
	if err != nil {
		return err
	}

	if exist == nil {
		return apierrors.ErrorNotFoundByUUID
	}

	item.UpdatedAt = s.timeController.Now()
	item.CreatedAt = exist.CreatedAt

	return s.repository.SaveAccount(item)
}
