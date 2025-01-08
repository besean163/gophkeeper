package route

import (
	"net/http"

	"github.com/besean163/gophkeeper/internal/server/api/dependencies"
	apierrors "github.com/besean163/gophkeeper/internal/server/api/errors"
)

func NotFoundRoute(dep dependencies.Dependencies) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dep.Logger.Debug("404 route")
		apierrors.SendError(w, http.StatusNotFound, "page not found")

	}
}
