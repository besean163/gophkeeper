package bucket

import models "github.com/besean163/gophkeeper/internal/models/server"

// DeleteCard удаление карты
func (s Service) DeleteCard(user models.User, uuid string) error {
	return s.repository.DeleteCard(uuid)
}
