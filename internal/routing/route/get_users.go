package route

import (
	"log"
	"net/http"
)

func GetUsersRoute() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("get users route")
	}
}
