package bucket

import "github.com/besean163/gophkeeper/internal/server/models"

// DeleteAccount удаление аккаунта
func (s Service) DeleteAccount(user models.User, uuid string) error {
	return s.repository.DeleteAccount(uuid)
}
