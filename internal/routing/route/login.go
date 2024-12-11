package route

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/besean163/gophkeeper/internal/auth"
)

func LoginRoute(s auth.AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("login route")
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

		user := s.GetUser(input.Login)
		if user == nil {
			sendError(w, http.StatusBadRequest, ErrorLoginUserNotExist.Error())
			return
		}

		tokenString, err := s.CreateUserToken(user)
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
