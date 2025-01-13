package api

import (
	models "github.com/besean163/gophkeeper/internal/models/client"
)

func (s Service) DeleteNote(user models.User, item models.Note, soft bool) error {
	var err error
	err = s.storeService.DeleteNote(user, item, soft)
	if err != nil {
		return err
	}

	err = s.syncer.Sync(user, SyncNodeNote)
	if err != nil {
		return err
	}

	return nil
}
