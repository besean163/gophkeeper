package api

import (
	"github.com/besean163/gophkeeper/internal/logger"
	models "github.com/besean163/gophkeeper/internal/models/client"
)

func (s Service) Sync(user models.User, nodes ...int) error {
	var err error
	for _, node := range nodes {
		switch node {
		case SyncNodeAccount:
			err = s.SyncAccounts(user)
		case SyncNodeNote:
			err = s.SyncNotes(user)
		case SyncNodeCard:
			err = s.SyncCards(user)
		}
		if err != nil {
			return err
		}
	}

	s.logger.Debug("sync all ...", logger.Field{Key: "user", Value: user.Login})
	return nil
}
