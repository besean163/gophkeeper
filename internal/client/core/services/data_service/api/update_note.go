package api

import (
	"github.com/besean163/gophkeeper/internal/client/core/models"
)

func (s Service) UpdateNote(user models.User, item models.Note) error {
	var err error
	err = s.storeService.UpdateNote(user, item)
	if err != nil {
		return err
	}
	err = s.syncer.SyncNotes(user)
	if err != nil {
		return err
	}

	return nil
}
