package middleware

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/besean163/gophkeeper/internal/models"
)

func AuthMiddleware() func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println("auth mid")

			body, err := io.ReadAll(r.Body)
			if err != nil {
				log.Println(err)
				return
			}
			log.Println(string(body))
			r.Body = io.NopCloser(bytes.NewBuffer(body))

			var test models.Test
			err = json.Unmarshal(body, &test)
			if err != nil {
				log.Println(err)
				return
			}
			log.Printf("Value of key: %s", test.Key)

			h.ServeHTTP(w, r)
		})
	}
}
