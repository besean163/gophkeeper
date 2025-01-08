package route

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"

	"github.com/besean163/gophkeeper/internal/logger"
	"github.com/besean163/gophkeeper/internal/server/api/dependencies"
	"github.com/besean163/gophkeeper/internal/server/api/entities"
	apierrors "github.com/besean163/gophkeeper/internal/server/api/errors"
)

func LoginRoute(dep dependencies.Dependencies) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		body, err := io.ReadAll(r.Body)
		if err != nil {
			apierrors.SendError(w, http.StatusBadRequest, apierrors.ErrorInvalidJSONData.Error())
			return
		}

		input := entities.LoginInput{}
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

		tokenString, err := dep.AuthService.LoginUser(input.Login, input.Password)
		if err != nil {
			if errors.Is(err, apierrors.ErrorUserNotExist) {
				apierrors.SendError(w, http.StatusBadRequest, apierrors.ErrorUserNotExist.Error())
				return
			} else {
				dep.Logger.Error("get token", logger.NewField("error", err.Error()))
				apierrors.SendError(w, http.StatusInternalServerError, apierrors.ErrorInternalUnknown.Error())
				return
			}
		}

		token := entities.TokenOutput{
			Token: tokenString,
		}

		result, err := json.Marshal(token)
		if err != nil {
			dep.Logger.Error("json make", logger.NewField("error", err.Error()))
			apierrors.SendError(w, http.StatusInternalServerError, apierrors.ErrorInternalUnknown.Error())
			return
		}

		w.Header().Add("Content-type", "application/json")
		w.Write(result)
	}
}
