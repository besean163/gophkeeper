// Package api пакет содержит методы и структуры для работы REST API
package api

import (
	"net/http"

	"github.com/besean163/gophkeeper/internal/server/api/middleware"
	"github.com/besean163/gophkeeper/internal/server/api/route"
	"github.com/besean163/gophkeeper/internal/server/dependencies"
	"github.com/go-chi/chi/v5"
)

// Handler обработчик запросов.
type Handler struct {
	handler http.Handler
}

// NewHandler создание обработчика запросов.
func NewHandler(dependencies dependencies.Dependencies) Handler {
	r := chi.NewRouter()
	r.Use(middleware.LogMiddleware(dependencies))
	r.NotFound(route.NotFoundRoute(dependencies))
	r.With(middleware.CheckContentTypeJSONMiddleware()).Post("/register", route.RegisterRoute(dependencies))
	r.With(middleware.CheckContentTypeJSONMiddleware()).Post("/login", route.LoginRoute(dependencies))
	r.Get("/ping", route.PingRoute())
	r.Route("/api/", func(r chi.Router) {
		r.Use(middleware.AuthMiddleware(dependencies))
		r.Route("/accounts", func(r chi.Router) {
			r.Get("/", route.AccountsRoute(dependencies))
			r.Post("/sync", route.AccountsSyncRoute(dependencies))
		})
		r.Route("/account", func(r chi.Router) {
			r.Post("/", route.AccountCreateRoute(dependencies))
			r.Put("/", route.AccountUpdateRoute(dependencies))
			r.Delete("/", route.AccountDeleteRoute(dependencies))
		})

		r.Route("/notes", func(r chi.Router) {
			r.Get("/", route.NotesRoute(dependencies))
			r.Post("/sync", route.NotesSyncRoute(dependencies))
		})
		r.Route("/note", func(r chi.Router) {
			r.Post("/", route.NoteCreateRoute(dependencies))
			r.Put("/", route.NoteUpdateRoute(dependencies))
			r.Delete("/", route.NoteDeleteRoute(dependencies))
		})

		r.Route("/cards", func(r chi.Router) {
			r.Get("/", route.CardsRoute(dependencies))
			r.Post("/sync", route.CardsSyncRoute(dependencies))
		})
		r.Route("/card", func(r chi.Router) {
			r.Post("/", route.CardCreateRoute(dependencies))
			r.Put("/", route.CardUpdateRoute(dependencies))
			r.Delete("/", route.CardDeleteRoute(dependencies))
		})
	})

	return Handler{
		handler: r,
	}
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.handler.ServeHTTP(w, r)
}
