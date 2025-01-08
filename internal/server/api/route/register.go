package route

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/besean163/gophkeeper/internal/logger"
	"github.com/besean163/gophkeeper/internal/server/api/dependencies"
	"github.com/besean163/gophkeeper/internal/server/api/entities"
	apierrors "github.com/besean163/gophkeeper/internal/server/api/errors"
)

func RegisterRoute(dep dependencies.Dependencies) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		body, err := io.ReadAll(r.Body)
		if err != nil {
			apierrors.SendError(w, http.StatusBadRequest, apierrors.ErrorInvalidJSONData.Error())
			return
		}

		input := entities.RegisterInput{}
		err = json.Unmarshal(body, &input)
		if err != nil {
			apierrors.SendError(w, http.StatusBadRequest, apierrors.ErrorInvalidJSONData.Error())
			return
		}

		validateError := input.Validate(http.StatusBadRequest)
		if validateError != nil {
			apierrors.SendError(w, validateError.Error.Code, validateError.Error.Description)
			return
		}

		tokenString, err := dep.AuthService.RegisterUser(input.Login, input.Password)
		if err != nil {
			if errors.Is(err, apierrors.ErrorUserExist) {
				apierrors.SendError(w, http.StatusBadRequest, err.Error())
			} else {
				dep.Logger.Error("get token", logger.NewField("error", err.Error()))
				apierrors.SendError(w, http.StatusInternalServerError, apierrors.ErrorInternalUnknown.Error())
			}
			return
		}

		token := entities.TokenOutput{
			Token: tokenString,
		}

		result, err := json.Marshal(token)
		if err != nil {
			log.Println("json make error:", err.Error())
			apierrors.SendError(w, http.StatusInternalServerError, apierrors.ErrorInternalUnknown.Error())
			return
		}

		w.Header().Add("Content-type", "application/json")
		w.Write(result)
	}
}
