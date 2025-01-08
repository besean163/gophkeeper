package route

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/besean163/gophkeeper/internal/logger"
	"github.com/besean163/gophkeeper/internal/server/api/dependencies"
	"github.com/besean163/gophkeeper/internal/server/api/entities"
	apierrors "github.com/besean163/gophkeeper/internal/server/api/errors"
	"github.com/besean163/gophkeeper/internal/server/models"
	ctxuser "github.com/besean163/gophkeeper/internal/server/utils/ctx_user"
)

func NotesSyncRoute(dep dependencies.Dependencies) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, ok := ctxuser.GetUserFromContext(r.Context())
		if !ok {
			dep.Logger.Error("get user", logger.NewField("error", "user not found in context"))
			apierrors.SendError(w, http.StatusUnauthorized, apierrors.ErrorNotAuthorized.Error())
			return
		}

		var err error
		body, err := io.ReadAll(r.Body)
		if err != nil {
			apierrors.SendError(w, http.StatusBadRequest, apierrors.ErrorInvalidJSONData.Error())
			return
		}

		input := entities.NotesSyncInput{}
		err = json.Unmarshal(body, &input)
		if err != nil {
			dep.Logger.Error("sync make json", logger.NewField("error", err.Error()))
			apierrors.SendError(w, http.StatusBadRequest, apierrors.ErrorInvalidJSONData.Error())
			return
		}

		externalNotes := make([]models.ExternalNote, 0)
		for _, note := range input.Notes {
			externalNote := models.ExternalNote{
				UUID:      note.UUID,
				Name:      note.Name,
				Content:   note.Content,
				CreatedAt: note.CreatedAt,
				UpdatedAt: note.UpdatedAt,
				DeletedAt: note.DeletedAt,
				SyncedAt:  note.SyncedAt,
			}
			externalNotes = append(externalNotes, externalNote)
		}

		err = dep.BucketService.SyncNotes(dep.BucketService, *user, externalNotes)
		if err != nil {
			dep.Logger.Error("sync notes", logger.NewField("error", err.Error()))
			apierrors.SendError(w, http.StatusInternalServerError, apierrors.ErrorInternalUnknown.Error())
			return
		}
	}
}
