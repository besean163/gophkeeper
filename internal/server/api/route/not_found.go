package route

import (
	"log"
	"net/http"

	apierrors "github.com/besean163/gophkeeper/internal/server/api/errors"
)

func NotFoundRoute() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("not found route")
		apierrors.SendError(w, http.StatusNotFound, "page not found")

	}
}
