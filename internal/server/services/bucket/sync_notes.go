package bucket

import (
	"github.com/besean163/gophkeeper/internal/server/interfaces"
	"github.com/besean163/gophkeeper/internal/server/models"
)

// SyncNotes синхронизация заметок
func (s Service) SyncNotes(service interfaces.BucketService, user models.User, externalItems []models.ExternalNote) error {
	items, err := s.repository.GetNotes(user)
	if err != nil {
		return err
	}

	created, updated, deleted := s.changeDetector.GetNotesChanges(user, items, externalItems)

	for _, item := range created {
		err := service.CreateNote(user, item)
		if err != nil {
			return err
		}
	}

	for _, item := range updated {
		err := service.UpdateNote(user, item)
		if err != nil {
			return err
		}
	}

	for _, uuid := range deleted {
		err := service.DeleteNote(user, uuid)
		if err != nil {
			return err
		}
	}
	return nil
}
