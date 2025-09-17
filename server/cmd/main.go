package main

import (
    "log"
    "sync"
	"github.com/nalindalal/CodeLookout/server/cmd/llm_cli"
	"github.com/nalindalal/CodeLookout/server/cmd/sonarqube_cli"
)

func main() {
    var wg sync.WaitGroup

    // Example: start a goroutine
    wg.Add(1)
    go func() {
        defer wg.Done()
        // Do some work here
        log.Println("Worker finished")
    }()

    // Wait for everything to shut down
    wg.Wait()
    log.Println("App shutdown complete.")
}

