package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Mentro-Org/CodeLookout/internal/api"
	"github.com/Mentro-Org/CodeLookout/internal/config"
	"github.com/joho/godotenv"
)

func main() {
	// Only load .env in local/dev environments
	if os.Getenv("APP_ENV") != "production" {
		_ = godotenv.Load()
	}

	cfg := config.Load()

	// Setup router
	r := api.NewRouter(cfg)

	// Start server
	addr := fmt.Sprintf(":%s", cfg.Port)
	fmt.Printf("Server is listening at http://localhost%s\n", addr)

	log.Fatal(http.ListenAndServe(addr, r))
}
