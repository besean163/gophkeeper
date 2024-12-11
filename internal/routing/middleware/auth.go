package middleware

import (
	"log"
	"net/http"
	"strings"

	jwttoken "github.com/besean163/gophkeeper/internal/jwt_token"
)

func AuthMiddleware(secret string) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println("auth mid")

			token := strings.Replace(r.Header.Get("Authorization"), "Bearer ", "", -1)
			log.Println(token)

			claims, err := jwttoken.GetClaimsByToken(secret, token)
			if err != nil {
				log.Println("get claim error")
				log.Println(err.Error())
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			userId, ok := claims["user_id"]
			if !ok {
				log.Println("user id not exist in claims")
				w.WriteHeader(http.StatusInternalServerError)
				return
			}

			log.Println("user find ", userId)
			h.ServeHTTP(w, r)
		})
	}
}
