package route

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/besean163/gophkeeper/internal/auth"
)

var (
	ErrorRegisterNotJSONData     = errors.New("expect JSON data")
	ErrorRegisterInvalidJSONData = errors.New("invalid JSON data")
	ErrorRegisterUserExist       = errors.New("user already exist")
)

func RegisterRoute(s auth.AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if jsonHeader := r.Header.Get("Content-type"); jsonHeader != "application/json" {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(ErrorRegisterNotJSONData.Error()))
			return
		}

		var err error
		// var user models.User
		input := struct {
			Login    string `json:"login"`
			Password string `json:"password"`
		}{}

		body, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(ErrorRegisterInvalidJSONData.Error()))
		}

		err = json.Unmarshal(body, &input)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
		}

		existUser := s.GetUser(input.Login)
		if existUser != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(ErrorRegisterUserExist.Error()))
		}

		/*
			принимает логин пароль отдает структуру с токеном или ошибку:
			- если пользователь существует
			- пароль недостаточен или содержит неверные символы
			- логин не содержит неверные символы
		*/
		log.Println("register route")
	}
}
