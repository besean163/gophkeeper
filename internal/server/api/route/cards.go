package route

import (
	"encoding/json"
	"net/http"

	"github.com/besean163/gophkeeper/internal/logger"
	"github.com/besean163/gophkeeper/internal/server/api/dependencies"
	"github.com/besean163/gophkeeper/internal/server/api/entities"
	apierrors "github.com/besean163/gophkeeper/internal/server/api/errors"
	ctxuser "github.com/besean163/gophkeeper/internal/server/utils/ctx_user"
)

func CardsRoute(dep dependencies.Dependencies) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, ok := ctxuser.GetUserFromContext(r.Context())
		if !ok {
			dep.Logger.Error("get user", logger.NewField("error", "user not found in context"))
			apierrors.SendError(w, http.StatusUnauthorized, apierrors.ErrorNotAuthorized.Error())
			return
		}

		cards, err := dep.BucketService.GetCards(*user)
		if err != nil {
			dep.Logger.Error("get cards", logger.NewField("error", err.Error()))
			apierrors.SendError(w, http.StatusInternalServerError, apierrors.ErrorInternalUnknown.Error())
			return
		}

		output := entities.GetCardsOutput{
			Cards: cards,
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
