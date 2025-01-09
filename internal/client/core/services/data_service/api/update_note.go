package api

import (
	models "github.com/besean163/gophkeeper/internal/models/client"
)

func (s Service) UpdateNote(user models.User, item models.Note) error {
	var err error
	err = s.storeService.UpdateNote(user, item)
	if err != nil {
		return err
	}
	err = s.syncer.Sync(user, SyncNodeNote)
	if err != nil {
		return err
	}

	return nil
}
