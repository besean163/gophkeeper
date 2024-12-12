package middleware

import (
	"errors"
	"net/http"

	apierrors "github.com/besean163/gophkeeper/internal/api/errors"
)

var (
	ErrorRegisterNotJSONData = errors.New("expect JSON data")
)

func CheckContentTypeJSONMiddleware() func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if h := r.Header.Get("Content-type"); h != "application/json" {
				apierrors.SendError(w, http.StatusBadRequest, ErrorRegisterNotJSONData.Error())
				return
			}
			h.ServeHTTP(w, r)
		})
	}
}
