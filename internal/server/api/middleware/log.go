package middleware

import (
	"net/http"

	"github.com/besean163/gophkeeper/internal/logger"
	"github.com/besean163/gophkeeper/internal/server/api/dependencies"
)

// LogMiddleware запись в жернал данных о запросе.
func LogMiddleware(dep dependencies.Dependencies) func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			dep.Logger.Debug("visit", logger.NewField("url", r.URL), logger.NewField("method", r.Method))
			h.ServeHTTP(w, r)
		})
	}
}
