package api

import (
	"github.com/besean163/gophkeeper/internal/logger"
	models "github.com/besean163/gophkeeper/internal/models/client"
	"github.com/besean163/gophkeeper/internal/server/api/entities/input"
)

func (s Service) syncNotesOnServer(user models.User) error {
	if !s.apiClient.HasConnection() {
		return nil
	}

	items, err := s.storeService.GetNotes(user)
	if err != nil {
		return err
	}

	apiItems := make([]input.NoteSync, 0)
	for _, a := range items {
		apiItem := input.NoteSync{
			UUID:      a.UUID,
			Name:      a.Name,
			Content:   a.Content,
			CreatedAt: a.CreatedAt,
			UpdatedAt: a.UpdatedAt,
			DeletedAt: a.DeletedAt,
			SyncedAt:  a.SyncedAt,
		}
		apiItems = append(apiItems, apiItem)
	}
	input := input.NotesSync{
		Notes: apiItems,
	}

	s.apiClient.SetToken(user.Token)
	err = s.apiClient.SyncNotes(input)
	if err != nil {
		return err
	}
	s.logger.Debug("sync notes on server ...", logger.Field{Key: "user", Value: user.Login})
	return nil
}
