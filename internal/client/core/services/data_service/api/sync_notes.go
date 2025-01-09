package api

import models "github.com/besean163/gophkeeper/internal/models/client"

func (s Service) SyncNotes(user models.User) error {
	var err error
	err = s.syncNotesOnServer(user)
	if err != nil {
		return err
	}
	err = s.syncNotesOnClient(user)
	if err != nil {
		return err
	}
	return nil
}
