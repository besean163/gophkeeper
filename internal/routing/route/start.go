package route

import (
	"log"
	"net/http"
)

func StartRoute() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("route")
		w.Write([]byte("Hi, i'm work"))
	}
}
