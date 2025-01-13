package bucket

import (
	models "github.com/besean163/gophkeeper/internal/models/server"
	apierrors "github.com/besean163/gophkeeper/internal/server/api/errors"
)

// UpdateCard обновление карты
func (s Service) UpdateCard(user models.User, item *models.Card) error {
	exist, err := s.repository.GetCard(item.UUID)
	if err != nil {
		return err
	}

	if exist == nil {
		return apierrors.ErrorNotFoundByUUID
	}

	item.UpdatedAt = s.timeController.Now()
	item.CreatedAt = exist.CreatedAt

	return s.repository.SaveCard(item)
}
