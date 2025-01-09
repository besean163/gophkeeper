package api

import (
	"github.com/besean163/gophkeeper/internal/logger"
	models "github.com/besean163/gophkeeper/internal/models/client"
)

func (s Service) SyncAll(user models.User) error {
	var err error
	err = s.SyncAccounts(user)
	if err != nil {
		return err
	}

	err = s.SyncNotes(user)
	if err != nil {
		return err
	}

	err = s.SyncCards(user)
	if err != nil {
		return err
	}

	s.logger.Debug("sync all ...", logger.Field{Key: "user", Value: user.Login})
	return nil
}
