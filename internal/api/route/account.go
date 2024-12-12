package route

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/besean163/gophkeeper/internal/api/entity"
	apierrors "github.com/besean163/gophkeeper/internal/api/errors"
	"github.com/besean163/gophkeeper/internal/bucket"
	"github.com/besean163/gophkeeper/internal/models"
)

func AccountRoute(s bucket.BucketService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("here")

		switch r.Method {
		case http.MethodPost:
			log.Println("account post route")

			var err error
			body, err := io.ReadAll(r.Body)
			if err != nil {
				apierrors.SendError(w, http.StatusBadRequest, apierrors.ErrorInvalidJSONData.Error())
				return
			}

			input := entity.PostAccountInput{}
			err = json.Unmarshal(body, &input)
			if err != nil {
				apierrors.SendError(w, http.StatusBadRequest, err.Error())
				return
			}

			account := models.Account{
				Name:     input.Name,
				Login:    input.Login,
				Password: input.Password,
			}

			if err := s.CreateAccount(&account); err != nil {
				apierrors.SendError(w, http.StatusBadRequest, err.Error())
				return
			}

		case http.MethodPut:
			log.Println("account put route")
			var err error
			body, err := io.ReadAll(r.Body)
			if err != nil {
				apierrors.SendError(w, http.StatusBadRequest, apierrors.ErrorInvalidJSONData.Error())
				return
			}

			input := entity.PutAccountInput{}
			err = json.Unmarshal(body, &input)
			if err != nil {
				apierrors.SendError(w, http.StatusBadRequest, err.Error())
				return
			}

			account := models.Account{
				ID:       input.ID,
				Name:     input.Name,
				Login:    input.Login,
				Password: input.Password,
			}

			if err := s.UpdateAccount(&account); err != nil {
				apierrors.SendError(w, http.StatusBadRequest, err.Error())
				return
			}
		case http.MethodDelete:
			log.Println("account delete route")
			var err error
			body, err := io.ReadAll(r.Body)
			if err != nil {
				apierrors.SendError(w, http.StatusBadRequest, apierrors.ErrorInvalidJSONData.Error())
				return
			}

			input := entity.DeleteAccountInput{}
			err = json.Unmarshal(body, &input)
			if err != nil {
				apierrors.SendError(w, http.StatusBadRequest, err.Error())
				return
			}

			if err := s.DeleteAccount(input.ID); err != nil {
				apierrors.SendError(w, http.StatusBadRequest, err.Error())
				return
			}
		}
	}
}
