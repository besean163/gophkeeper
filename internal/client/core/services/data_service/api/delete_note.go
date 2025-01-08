package api

import (
	"github.com/besean163/gophkeeper/internal/client/core/models"
)

func (s Service) DeleteNote(user models.User, item models.Note, soft bool) error {
	var err error
	err = s.storeService.DeleteNote(user, item, soft)
	if err != nil {
		return err
	}

	err = s.syncer.SyncNotes(user)
	if err != nil {
		return err
	}

	return nil
}
