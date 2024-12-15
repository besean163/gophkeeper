package route

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/besean163/gophkeeper/internal/server/api/entity"
	apierrors "github.com/besean163/gophkeeper/internal/server/api/errors"
	"github.com/besean163/gophkeeper/internal/server/interfaces"
)

func AccountsRoute(s interfaces.BucketService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("get accounts route")

		accounts := s.GetAccounts()
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
