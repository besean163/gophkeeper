package route

import (
	"log"
	"net/http"

	"github.com/besean163/gophkeeper/internal/bucket"
)

func GetAccountsRoute(s bucket.BucketService) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("get accounts route")
	}
}
