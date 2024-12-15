package middleware

import (
	"errors"
	"log"
	"net/http"
	"strings"

	apierrors "github.com/besean163/gophkeeper/internal/server/api/errors"
	jwttoken "github.com/besean163/gophkeeper/internal/server/utils/jwt_token"
)

func AuthMiddleware(secret string) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println("auth mid")
			log.Println(r.Header.Get("Authorization"))

			token := strings.Replace(r.Header.Get("Authorization"), "Bearer ", "", -1)
			log.Println(token)

			claims, err := jwttoken.GetClaimsByToken(secret, token)
			if err != nil {
				log.Println("get claim error")
				log.Println(err.Error())
				apierrors.SendError(w, http.StatusInternalServerError, err.Error())
				return
			}

			userId, ok := claims["user_id"]
			if !ok {
				err := errors.New("user id not exist in claims")
				log.Println(err.Error())
				apierrors.SendError(w, http.StatusInternalServerError, err.Error())
				return
			}

			log.Println("user find ", userId)
			h.ServeHTTP(w, r)
		})
	}
}
