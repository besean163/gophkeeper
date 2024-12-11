package routing

import (
	"net/http"

	"github.com/besean163/gophkeeper/internal/auth"
	"github.com/besean163/gophkeeper/internal/bucket"
	"github.com/besean163/gophkeeper/internal/routing/middleware"
	"github.com/besean163/gophkeeper/internal/routing/route"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	handler http.Handler
	auth.AuthService
}

func NewHandler(secret string, authService auth.AuthService, bucketService bucket.BucketService) Handler {
	r := chi.NewRouter()
	r.Use(middleware.LogMiddleware())
	r.HandleFunc("/", route.StartRoute())
	r.Post("/register", route.RegisterRoute(authService))
	r.Post("/login", route.LoginRoute(authService))
	r.Route("/api/", func(r chi.Router) {
		r.Use(middleware.AuthMiddleware(secret))
		r.Get("/accounts", route.GetAccountsRoute(bucketService))
	})

	return Handler{
		handler:     r,
		AuthService: authService,
	}
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.handler.ServeHTTP(w, r)
}
