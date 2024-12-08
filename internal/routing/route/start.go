package route

import (
	"log"
	"net/http"
)

func StartRoute() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("start route")
	}
}
