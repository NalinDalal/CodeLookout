package handlers

import (
	"net/http"

	"github.com/Mentro-Org/CodeLookout/internal/config"
	"github.com/Mentro-Org/CodeLookout/internal/core"
	ghclient "github.com/Mentro-Org/CodeLookout/internal/github"
	pr "github.com/Mentro-Org/CodeLookout/internal/handlers/pullrequest"
	"github.com/Mentro-Org/CodeLookout/internal/llm"

	"github.com/google/go-github/github"
)

func WebhookHandler(cfg *config.Config, ghClientFactory *ghclient.ClientFactory, aiClient llm.AIClient) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		payload, err := github.ValidatePayload(r, []byte(cfg.WebhookSecret))
		if err != nil {
			http.Error(w, "Invalid payload signature", http.StatusUnauthorized)
			return
		}

		event, err := github.ParseWebHook(github.WebHookType(r), payload)
		if err != nil {
			http.Error(w, "Could not parse webhook", http.StatusBadRequest)
			return
		}

		switch e := event.(type) {
		case *github.PullRequestEvent:
			handler := routePullRequestEvent(e.GetAction(), cfg, ghClientFactory, aiClient)
			if handler == nil {
				http.Error(w, "Unsupported PR action", http.StatusNotImplemented)
				return
			}
			if err := handler.Handle(ctx, e); err != nil {
				http.Error(w, "Handler error: "+err.Error(), http.StatusInternalServerError)
				return
			}
		default:
			http.Error(w, "Unsupported event type", http.StatusNotImplemented)
			return

		}

		w.WriteHeader(http.StatusOK)
	}
}

func routePullRequestEvent(action string, cfg *config.Config, ghClientFactory *ghclient.ClientFactory, aiClient llm.AIClient) core.PullRequestHandler {
	switch action {
	case "opened":
		return &pr.PullRequestOpenedHandler{Cfg: cfg, GHClientFactory: ghClientFactory, AIClient: aiClient}
	case "edited":
		return &pr.PullRequestEditedHandler{Cfg: cfg, GHClientFactory: ghClientFactory, AIClient: aiClient}
	case "synchronize":
		return &pr.PullRequestSynchronizedHandler{Cfg: cfg, GHClientFactory: ghClientFactory, AIClient: aiClient}
	default:
		return nil
	}
}
