package route

import (
	"encoding/json"
	"net/http"

	"github.com/besean163/gophkeeper/internal/logger"
	"github.com/besean163/gophkeeper/internal/server/api/entities/output"
	apierrors "github.com/besean163/gophkeeper/internal/server/api/errors"
	"github.com/besean163/gophkeeper/internal/server/dependencies"
	ctxuser "github.com/besean163/gophkeeper/internal/server/utils/ctx_user"
)

func NotesRoute(dep dependencies.Dependencies) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, ok := ctxuser.GetUserFromContext(r.Context())
		if !ok {
			dep.Logger.Error("get user", logger.NewField("error", "user not found in context"))
			apierrors.SendError(w, http.StatusUnauthorized, apierrors.ErrorNotAuthorized.Error())
			return
		}

		notes, err := dep.BucketService.GetNotes(*user)
		if err != nil {
			dep.Logger.Error("get notes", logger.NewField("error", err.Error()))
			apierrors.SendError(w, http.StatusInternalServerError, apierrors.ErrorInternalUnknown.Error())
			return
		}

		output := output.GetNotes{
			Notes: notes,
		}

		result, err := json.Marshal(output)
		if err != nil {
			dep.Logger.Error("json make", logger.NewField("error", err.Error()))
			apierrors.SendError(w, http.StatusInternalServerError, apierrors.ErrorInternalUnknown.Error())
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.Write(result)
	}
}
