package middleware

import (
	"log"
	"net/http"
)

func LogMiddleware() func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println("log mid")
			h.ServeHTTP(w, r)
		})
	}
}
