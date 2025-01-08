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

func CardCreateRoute(dep dependencies.Dependencies) http.HandlerFunc {
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

		input := entities.CardCreateInput{}
		err = json.Unmarshal(body, &input)
		if err != nil {
			dep.Logger.Error("post make json", logger.NewField("error", err.Error()))
			apierrors.SendError(w, http.StatusBadRequest, apierrors.ErrorInvalidJSONData.Error())
			return
		}

		card := models.Card{
			UserID: user.ID,
			Name:   input.Name,
			Number: input.Number,
			Exp:    input.Exp,
			CVV:    input.CVV,
		}

		if err := dep.BucketService.CreateCard(*user, &card); err != nil {
			apierrors.SendError(w, http.StatusBadRequest, err.Error())
			return
		}
	}
}
