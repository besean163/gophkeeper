package middleware

import (
	"context"
	"errors"
	"log"
	"net/http"
	"reflect"
	"strings"

	"github.com/besean163/gophkeeper/internal/client/tui/logger"
	"github.com/besean163/gophkeeper/internal/server/api/entity"
	apierrors "github.com/besean163/gophkeeper/internal/server/api/errors"
	"github.com/besean163/gophkeeper/internal/server/interfaces"
	jwttoken "github.com/besean163/gophkeeper/internal/server/utils/jwt_token"
)

func AuthMiddleware(authService interfaces.AuthService, secret string) func(h http.Handler) http.Handler {
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

			logger.Get().Println(claims)
			userId, ok := claims["user_id"].(float64)
			logger.Get().Println(ok)
			logger.Get().Println(userId)
			logger.Get().Println(reflect.TypeOf(userId))
			if !ok {
				err := errors.New("user id not exist in claims")
				log.Println(err.Error())
				apierrors.SendError(w, http.StatusInternalServerError, err.Error())
				return
			}

			user, err := authService.GetUser(int(userId))
			if err != nil {
				log.Println("get user error")
				log.Println(err.Error())
				apierrors.SendError(w, http.StatusInternalServerError, err.Error())
				return
			}

			ctx := context.WithValue(r.Context(), entity.RequestUserKey("user"), user)

			log.Println("user find ", userId)
			h.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
