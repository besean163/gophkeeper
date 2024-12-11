package route

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/besean163/gophkeeper/internal/auth"
	jwttoken "github.com/besean163/gophkeeper/internal/jwt_token"
	"github.com/golang-jwt/jwt/v5"
)

var (
	ErrorRegisterNotJSONData     = errors.New("expect JSON data")
	ErrorRegisterInvalidJSONData = errors.New("invalid JSON data")
	ErrorRegisterUserExist       = errors.New("user already exist")
)

func RegisterRoute(secret string, s auth.AuthService) http.HandlerFunc {
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

		claims := jwt.MapClaims{
			"user_id": 123,
			"exp":     time.Now().Add(1 * time.Second).Unix(), // Время истечения
		}
		tokenString, err := jwttoken.GetJWTToken(secret, claims)
		if err != nil {
			log.Println("get token error:", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		type Token struct {
			Token string `json:"token"`
		}
		token := Token{
			Token: tokenString,
		}

		result, err := json.Marshal(token)
		if err != nil {
			log.Println("json make error:", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(result)

		/*
			принимает логин пароль отдает структуру с токеном или ошибку:
			- если пользователь существует
			- пароль недостаточен или содержит неверные символы
			- логин не содержит неверные символы
		*/
		log.Println("register route")
	}
}
