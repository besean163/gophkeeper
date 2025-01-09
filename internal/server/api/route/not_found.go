package route

import (
	"net/http"

	apierrors "github.com/besean163/gophkeeper/internal/server/api/errors"
	"github.com/besean163/gophkeeper/internal/server/dependencies"
)

func NotFoundRoute(dep dependencies.Dependencies) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dep.Logger.Debug("404 route")
		apierrors.SendError(w, http.StatusNotFound, "page not found")

	}
}
