package route

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/besean163/gophkeeper/internal/logger"
	"github.com/besean163/gophkeeper/internal/server/api/entity"
	apierrors "github.com/besean163/gophkeeper/internal/server/api/errors"
	"github.com/besean163/gophkeeper/internal/server/interfaces"
	ctxuser "github.com/besean163/gophkeeper/internal/server/utils/ctx_user"
)

func AccountsRoute(s interfaces.BucketService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		logger.Get().Println(r.Context())
		user, ok := ctxuser.GetUserFromContext(r.Context())
		if !ok {
			http.Error(w, "User not found in context", http.StatusUnauthorized)
			return
		}

		log.Println("get accounts route")

		accounts, err := s.GetAccounts(*user)
		if err != nil {
			log.Println("get accounts error:", err.Error())
			apierrors.SendError(w, http.StatusInternalServerError, apierrors.ErrorInternalUnknown.Error())
			return
		}
		output := entity.GetAccountsOutput{
			Accounts: accounts,
		}

		result, err := json.Marshal(output)
		if err != nil {
			log.Println("json make error:", err.Error())
			apierrors.SendError(w, http.StatusInternalServerError, apierrors.ErrorInternalUnknown.Error())
			return
		}

		w.Header().Add("Content-Type", "application/json")
		w.Write(result)
	}
}
