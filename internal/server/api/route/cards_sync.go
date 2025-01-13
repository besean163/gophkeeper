package route

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/besean163/gophkeeper/internal/logger"
	clientmodels "github.com/besean163/gophkeeper/internal/models/client"

	"github.com/besean163/gophkeeper/internal/server/api/entities/input"
	apierrors "github.com/besean163/gophkeeper/internal/server/api/errors"
	"github.com/besean163/gophkeeper/internal/server/dependencies"

	ctxuser "github.com/besean163/gophkeeper/internal/server/utils/ctx_user"
)

func CardsSyncRoute(dep dependencies.Dependencies) http.HandlerFunc {
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

		input := input.CardsSync{}
		err = json.Unmarshal(body, &input)
		if err != nil {
			dep.Logger.Error("sync make json", logger.NewField("error", err.Error()))
			apierrors.SendError(w, http.StatusBadRequest, apierrors.ErrorInvalidJSONData.Error())
			return
		}

		externalCards := make([]clientmodels.Card, 0)
		for _, card := range input.Cards {
			externalCard := clientmodels.Card{
				UUID:      card.UUID,
				Name:      card.Name,
				Number:    card.Number,
				Exp:       card.Exp,
				CVV:       card.CVV,
				CreatedAt: card.CreatedAt,
				UpdatedAt: card.UpdatedAt,
				DeletedAt: card.DeletedAt,
				SyncedAt:  card.SyncedAt,
			}
			externalCards = append(externalCards, externalCard)
		}

		err = dep.BucketService.SyncCards(dep.BucketService, *user, externalCards)
		if err != nil {
			dep.Logger.Error("sync cards", logger.NewField("error", err.Error()))
			apierrors.SendError(w, http.StatusInternalServerError, apierrors.ErrorInternalUnknown.Error())
			return
		}
	}
}
