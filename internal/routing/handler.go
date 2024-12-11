package routing

import (
	"net/http"

	"github.com/besean163/gophkeeper/internal/auth"
	"github.com/besean163/gophkeeper/internal/routing/middleware"
	"github.com/besean163/gophkeeper/internal/routing/route"
	"github.com/go-chi/chi/v5"
)

type Handler struct {
	handler http.Handler
	auth.AuthService
}

func NewHandler(secret string, authService auth.AuthService) Handler {
	r := chi.NewRouter()
	r.Use(middleware.AuthMiddleware(), middleware.LogMiddleware())
	r.HandleFunc("/", route.StartRoute())
	r.Route("/api/", func(r chi.Router) {
		r.Route("/user/", func(r chi.Router) {
			r.Post("/login", route.LoginRoute(authService))
			r.Post("/register", route.RegisterRoute(secret, authService))
		})
	})

	return Handler{
		handler:     r,
		AuthService: authService,
	}
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.handler.ServeHTTP(w, r)
}
