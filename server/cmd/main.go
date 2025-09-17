package main

import (
    "log"
    "sync"
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

