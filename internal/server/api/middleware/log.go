package middleware

import (
	"log"
	"net/http"
)

func LogMiddleware() func(h http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println("log mid")

			// body, err := io.ReadAll(r.Body)
			// if err != nil {
			// 	log.Println(err)
			// 	return
			// }
			// log.Println(string(body))

			// r.Body = io.NopCloser(bytes.NewBuffer(body))
			h.ServeHTTP(w, r)
		})
	}
}
