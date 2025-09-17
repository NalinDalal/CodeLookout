package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"context"

	"github.com/Mentro-Org/CodeLookout/internal/core"
	db "github.com/Mentro-Org/CodeLookout/internal/db"
)

type AnalyticsResponse struct {
	Results []db.LLMAnalytics `json:"results"`
}

// GET /api/analytics?limit=50&offset=0
func HandleLLMAnalytics(appDeps *core.AppDeps) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
	       limit := 50
	       offset := 0
	       if l := r.URL.Query().Get("limit"); l != "" {
		       if v, err := strconv.Atoi(l); err == nil {
			       limit = v
		       }
	       }
	       if o := r.URL.Query().Get("offset"); o != "" {
		       if v, err := strconv.Atoi(o); err == nil {
			       offset = v
		       }
	       }
	       // Filters
	       filters := db.LLMAnalyticsFilters{
		       Error:   r.URL.Query().Get("error"),
		       Repo:    r.URL.Query().Get("repo"),
		       Owner:   r.URL.Query().Get("owner"),
		       PRNumber: r.URL.Query().Get("pr_number"),
		       Since:   r.URL.Query().Get("since"),
		       Until:   r.URL.Query().Get("until"),
	       }
	       results, err := db.ListLLMAnalyticsFiltered(ctx, appDeps.DBPool, limit, offset, filters)
	       if err != nil {
		       w.WriteHeader(http.StatusInternalServerError)
		       w.Write([]byte("error fetching analytics"))
		       return
	       }
	       w.Header().Set("Content-Type", "application/json")
	       json.NewEncoder(w).Encode(AnalyticsResponse{Results: results})
	}
}
