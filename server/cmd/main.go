package main
	}()

	// Wait for everything to shut down
	wg.Wait()
	log.Println("App shutdown complete.")
}
