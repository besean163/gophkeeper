package api

import (
	"github.com/besean163/gophkeeper/internal/client/core/services/data_service/api/changes"
	"github.com/besean163/gophkeeper/internal/logger"
	models "github.com/besean163/gophkeeper/internal/models/client"
	servermodels "github.com/besean163/gophkeeper/internal/models/server"
)

func (s Service) syncNotesOnClient(user models.User) error {
	if !s.apiClient.HasConnection() {
		return nil
	}

	items, err := s.storeService.GetNotes(user)
	if err != nil {
		return err
	}

	s.apiClient.SetToken(user.Token)
	apiOutput, err := s.apiClient.GetNotes()
	if err != nil {
		return err
	}

	externalItems := make([]servermodels.Note, 0)
	for _, apiItem := range apiOutput.Notes {
		externalItem := servermodels.Note{
			UUID:      apiItem.UUID,
			UserID:    user.ID,
			Name:      apiItem.Name,
			Content:   apiItem.Content,
			CreatedAt: apiItem.CreatedAt,
			UpdatedAt: apiItem.UpdatedAt,
		}
		externalItems = append(externalItems, externalItem)
	}

	compare := changes.NoteCompare{
		Items:        items,
		CompareItems: externalItems,
	}
	changes := s.changeDetector.GetNoteChanges(user, compare)

	for _, item := range changes.Created {
		err := s.storeService.CreateNote(user, item)
		if err != nil {
			return err
		}
	}

	for _, item := range changes.Updated {
		err := s.storeService.UpdateNote(user, item)
		if err != nil {
			return err
		}
	}

	for _, item := range changes.Deleted {
		err := s.storeService.DeleteNote(user, item, false)
		if err != nil {
			return err
		}
	}

	s.logger.Debug("sync notes on client ...", logger.Field{Key: "user", Value: user.Login})
	return nil
}
