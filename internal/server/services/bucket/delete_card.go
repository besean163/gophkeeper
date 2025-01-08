package bucket

import "github.com/besean163/gophkeeper/internal/server/models"

// DeleteCard удаление карты
func (s Service) DeleteCard(user models.User, uuid string) error {
	return s.repository.DeleteCard(uuid)
}
