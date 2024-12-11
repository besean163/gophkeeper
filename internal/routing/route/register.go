package route

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/besean163/gophkeeper/internal/auth"
	apierrors "github.com/besean163/gophkeeper/internal/errors/api_errors"
	"github.com/besean163/gophkeeper/internal/models"
)

func RegisterRoute(s auth.AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("register route")
		var err error
		if jsonHeader := r.Header.Get("Content-type"); jsonHeader != "application/json" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(ErrorRegisterNotJSONData.Error()))
			return
		}

		w.Header().Add("Content-type", "application/json")
		body, err := io.ReadAll(r.Body)
		if err != nil {
			sendError(w, http.StatusBadRequest, ErrorRegisterInvalidJSONData.Error())
			return
		}

		input := input{}
		err = json.Unmarshal(body, &input)
		if err != nil {
			sendError(w, http.StatusBadRequest, err.Error())
			return
		}

		validateError := input.validate(http.StatusBadRequest)
		if validateError != nil {
			sendError(w, validateError.Error.Code, validateError.Error.Description)
			return
		}

		existUser := s.GetUser(input.Login)
		if existUser != nil {
			sendError(w, http.StatusBadRequest, ErrorRegisterUserExist.Error())
			return
		}

		user := models.User{
			Login:    input.Login,
			Password: input.Password,
		}

		tokenString, err := s.RegisterUser(&user)
		if err != nil {
			log.Println("get token error:", err.Error())
			sendError(w, http.StatusInternalServerError, ErrorInternalUnknown.Error())
			return
		}

		token := output{
			Token: tokenString,
		}

		result, err := json.Marshal(token)
		if err != nil {
			log.Println("json make error:", err.Error())
			sendError(w, http.StatusInternalServerError, ErrorInternalUnknown.Error())
			return
		}

		w.Write(result)
	}
}

func makeError(code int, description string) *apierrors.Error {
	e := &apierrors.Error{}
	e.Error.Code = code
	e.Error.Description = description
	return e
}

func sendError(w http.ResponseWriter, code int, description string) {
	apiError := apierrors.Error{}
	apiError.Error.Code = code
	apiError.Error.Description = description

	errorBody, err := json.Marshal(apiError)
	if err != nil {
		log.Println("send error fail")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(apiError.Error.Code)
	w.Write(errorBody)
}
