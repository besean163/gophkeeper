package routing

import (
	"net/http"

	"github.com/besean163/gophkeeper/internal/routing/middleware"
	"github.com/besean163/gophkeeper/internal/routing/route"
	"github.com/go-chi/chi/v5"
)

func NewHandler() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.AuthMiddleware(), middleware.LogMiddleware())
	r.Get("/", route.StartRoute())

	return r
}
