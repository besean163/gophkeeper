package bucket

import (
	models "github.com/besean163/gophkeeper/internal/models/server"
	apierrors "github.com/besean163/gophkeeper/internal/server/api/errors"
)

// UpdateNote обновление заметки
func (s Service) UpdateNote(user models.User, item *models.Note) error {
	exist, err := s.repository.GetNote(item.UUID)
	if err != nil {
		return err
	}

	if exist == nil {
		return apierrors.ErrorNotFoundByUUID
	}

	item.UpdatedAt = s.timeController.Now()
	item.CreatedAt = exist.CreatedAt

	return s.repository.SaveNote(item)
}
