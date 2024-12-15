package api

import (
	"net/http"

	"github.com/besean163/gophkeeper/internal/server/api/middleware"
	"github.com/besean163/gophkeeper/internal/server/api/route"
	"github.com/besean163/gophkeeper/internal/server/interfaces"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	handler http.Handler
	interfaces.AuthService
}

func NewHandler(secret string, authService interfaces.AuthService, bucketService interfaces.BucketService) Handler {
	r := chi.NewRouter()
	r.Use(middleware.LogMiddleware())
	r.NotFound(route.NotFoundRoute())
	r.With(middleware.CheckContentTypeJSONMiddleware()).Post("/register", route.RegisterRoute(authService))
	r.With(middleware.CheckContentTypeJSONMiddleware()).Post("/login", route.LoginRoute(authService))
	r.Route("/api/", func(r chi.Router) {
		r.Use(middleware.AuthMiddleware(secret))
		r.Get("/accounts", route.AccountsRoute(bucketService))
		// r.Get("/account", nil)
		r.Post("/account", route.AccountRoute(bucketService))
		r.Put("/account", route.AccountRoute(bucketService))
		r.Delete("/account", route.AccountRoute(bucketService))
	})

	return Handler{
		handler:     r,
		AuthService: authService,
	}
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.handler.ServeHTTP(w, r)
}
