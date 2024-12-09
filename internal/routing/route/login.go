package route

import (
	"log"
	"net/http"

	"github.com/besean163/gophkeeper/internal/auth"
)

func LoginRoute(s auth.AuthService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("login route")
		/*
			принимает логин пароль отдает структуру с токеном или ошибку:
			- если пользователя нет
		*/
	}
}
