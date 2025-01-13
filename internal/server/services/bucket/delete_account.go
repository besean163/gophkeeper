package bucket

import models "github.com/besean163/gophkeeper/internal/models/server"

// DeleteAccount удаление аккаунта
func (s Service) DeleteAccount(user models.User, uuid string) error {
	return s.repository.DeleteAccount(uuid)
}
