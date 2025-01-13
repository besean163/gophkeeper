package bucket

import (
	clientmodels "github.com/besean163/gophkeeper/internal/models/client"
	models "github.com/besean163/gophkeeper/internal/models/server"

	"github.com/besean163/gophkeeper/internal/server/interfaces"
	"github.com/besean163/gophkeeper/internal/server/services/bucket/changes"
)

// SyncNotes синхронизация заметок
func (s Service) SyncNotes(service interfaces.BucketService, user models.User, externalItems []clientmodels.Note) error {
	items, err := s.repository.GetNotes(user)
	if err != nil {
		return err
	}

	compare := changes.NoteCompare{
		Items:        items,
		CompareItems: externalItems,
	}
	changes := s.changeDetector.GetNotesChanges(user, compare)

	for _, item := range changes.Created {
		err := service.CreateNote(user, item)
		if err != nil {
			return err
		}
	}

	for _, item := range changes.Updated {
		err := service.UpdateNote(user, item)
		if err != nil {
			return err
		}
	}

	for _, uuid := range changes.Deleted {
		err := service.DeleteNote(user, uuid)
		if err != nil {
			return err
		}
	}
	return nil
}
