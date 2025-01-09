package route

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/besean163/gophkeeper/internal/logger"
	models "github.com/besean163/gophkeeper/internal/models/server"
	"github.com/besean163/gophkeeper/internal/server/api/entities/input"
	apierrors "github.com/besean163/gophkeeper/internal/server/api/errors"
	"github.com/besean163/gophkeeper/internal/server/dependencies"

	ctxuser "github.com/besean163/gophkeeper/internal/server/utils/ctx_user"
)

func AccountUpdateRoute(dep dependencies.Dependencies) http.HandlerFunc {
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

		input := input.AccountUpdate{}
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

		account := models.Account{
			UUID:     input.UUID,
			UserID:   user.ID,
			Name:     input.Name,
			Login:    input.Login,
			Password: input.Password,
		}

		if err := dep.BucketService.UpdateAccount(*user, &account); err != nil {
			apierrors.SendError(w, http.StatusBadRequest, err.Error())
			return
		}
	}
}
