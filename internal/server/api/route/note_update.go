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

func NoteUpdateRoute(dep dependencies.Dependencies) http.HandlerFunc {
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

		input := entities.NoteUpdateInput{}
		err = json.Unmarshal(body, &input)
		if err != nil {
			dep.Logger.Error("put make json", logger.NewField("error", err.Error()))
			apierrors.SendError(w, http.StatusBadRequest, apierrors.ErrorInvalidJSONData.Error())
			return
		}

		validateError := input.Validate(http.StatusBadRequest)
		if validateError != nil {
			apierrors.SendError(w, validateError.Error.Code, validateError.Error.Description)
			return
		}

		note := models.Note{
			UUID:    input.UUID,
			UserID:  user.ID,
			Name:    input.Name,
			Content: input.Content,
		}

		if err := dep.BucketService.UpdateNote(*user, &note); err != nil {
			apierrors.SendError(w, http.StatusBadRequest, err.Error())
			return
		}
	}
}
