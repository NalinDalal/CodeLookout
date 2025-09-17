package api

import (
    "net/http"
    "github.com/nalindalal/CodeLookout/server/internal/handlers"
    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
    "github.com/nalindalal/CodeLookout/server/internal/core"
)

func NewRouter(appDeps *core.AppDeps) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

       r.Route("/api", func(r chi.Router) {
	       r.Get("/health-check", func(w http.ResponseWriter, r *http.Request) {
		       w.WriteHeader(http.StatusOK)
	       })
	       r.Post("/webhook", handlers.HandleWebhook(appDeps))
	       r.Get("/analytics", HandleLLMAnalytics(appDeps))
       })

       // Serve analytics dashboard UI
       r.Get("/analytics", func(w http.ResponseWriter, r *http.Request) {
	       http.ServeFile(w, r, "internal/analyticsui/index.html")
       })
	return r
}
