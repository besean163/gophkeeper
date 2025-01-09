package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/besean163/gophkeeper/internal/logger"
	"github.com/besean163/gophkeeper/internal/server/api/entities"
	apierrors "github.com/besean163/gophkeeper/internal/server/api/errors"
	"github.com/besean163/gophkeeper/internal/server/dependencies"
)

// AuthMiddleware авторизации пользователя.
func AuthMiddleware(dep dependencies.Dependencies) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := strings.Replace(r.Header.Get("Authorization"), "Bearer ", "", -1)

			userId, err := dep.Tokener.GetUserId(token)
			if err != nil {
				dep.Logger.Error("get user id", logger.NewField("error", err.Error()))
				apierrors.SendError(w, http.StatusInternalServerError, err.Error())
				return
			}

			user, err := dep.AuthService.GetUser(userId)
			if err != nil {
				dep.Logger.Error("get user", logger.NewField("error", err.Error()))
				apierrors.SendError(w, http.StatusInternalServerError, err.Error())
				return
			}

			ctx := context.WithValue(r.Context(), entities.RequestUserKey("user"), user)
			h.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
