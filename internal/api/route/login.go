package route

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/besean163/gophkeeper/internal/api/entity"
	apierrors "github.com/besean163/gophkeeper/internal/api/errors"
	"github.com/besean163/gophkeeper/internal/auth"
)

func LoginRoute(s auth.AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("login route")

		var err error
		body, err := io.ReadAll(r.Body)
		if err != nil {
			apierrors.SendError(w, http.StatusBadRequest, apierrors.ErrorInvalidJSONData.Error())
			return
		}

		input := entity.LoginInput{}
		err = json.Unmarshal(body, &input)
		if err != nil {
			apierrors.SendError(w, http.StatusBadRequest, err.Error())
			return
		}

		validateError := input.Validate(http.StatusBadRequest)
		if validateError != nil {
			apierrors.SendError(w, validateError.Error.Code, validateError.Error.Description)
			return
		}

		user := s.GetUser(input.Login)
		if user == nil {
			apierrors.SendError(w, http.StatusBadRequest, apierrors.ErrorLoginUserNotExist.Error())
			return
		}

		tokenString, err := s.CreateUserToken(user)
		if err != nil {
			log.Println("get token error:", err.Error())
			apierrors.SendError(w, http.StatusInternalServerError, apierrors.ErrorInternalUnknown.Error())
			return
		}

		token := entity.TokenOutput{
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
