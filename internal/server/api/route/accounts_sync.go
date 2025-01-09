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

func AccountsSyncRoute(dep dependencies.Dependencies) http.HandlerFunc {
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

		input := input.AccountsSync{}
		err = json.Unmarshal(body, &input)
		if err != nil {
			dep.Logger.Error("sync make json", logger.NewField("error", err.Error()))
			apierrors.SendError(w, http.StatusBadRequest, apierrors.ErrorInvalidJSONData.Error())
			return
		}

		extAccounts := make([]clientmodels.Account, 0)
		for _, a := range input.Accounts {
			extAccount := clientmodels.Account{
				UUID:      a.UUID,
				Name:      a.Name,
				Login:     a.Login,
				Password:  a.Password,
				CreatedAt: a.CreatedAt,
				UpdatedAt: a.UpdatedAt,
				DeletedAt: a.DeletedAt,
				SyncedAt:  a.SyncedAt,
			}
			extAccounts = append(extAccounts, extAccount)
		}

		err = dep.BucketService.SyncAccounts(dep.BucketService, *user, extAccounts)
		if err != nil {
			dep.Logger.Error("sync accounts", logger.NewField("error", err.Error()))
			apierrors.SendError(w, http.StatusInternalServerError, apierrors.ErrorInternalUnknown.Error())
			return
		}
	}
}
