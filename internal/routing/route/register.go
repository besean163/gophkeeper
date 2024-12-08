package route

import (
	"log"
	"net/http"
)

func RegisterRoute() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		/*
			принимает логин пароль отдает структуру с токеном или ошибку:
			- если пользователь существует
			- пароль недостаточен или содержит неверные символы
			- логин не содержит неверные символы
		*/
		log.Println("register route")
	}
}
