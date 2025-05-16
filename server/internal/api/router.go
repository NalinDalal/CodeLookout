package api

import (
	"net/http"

	"github.com/Mentro-Org/CodeLookout/internal/config"
	"github.com/Mentro-Org/CodeLookout/internal/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewRouter(cfg *config.Config) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Route("/api", func(r chi.Router) {
		r.Post("/webhook", handlers.WebhookHandler(cfg))
	})

	return r
}
