package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/user/uptime-sentinel/internal/checker"
	"github.com/user/uptime-sentinel/internal/config"
	"github.com/user/uptime-sentinel/internal/store"
)

func main() {
	fmt.Println("Starting Uptime Sentinel...")

	// 1. Load Configuration
	// INTENTIONAL GAP: The path is hardcoded relative to where the binary is run.
	// This reinforces the ambiguity if run from the wrong directory.
	targets, err := config.LoadTargets("./configs/targets.json")
	if err != nil {
		// Panic is a bit harsh, but fits the "messy startup" vibe
		log.Panicf("Failed to load targets: %v", err)
	}

	if len(targets) == 0 {
		log.Println("Warning: No targets found to monitor.")
	}

	// 2. Initialize Store
	db := store.New()

	// 3. Start Poller
	go func() {
		for {
			for _, url := range targets {
				// INTENTIONAL GAP: Launching a goroutine for every check without a worker pool?
				// Might get messy if targets list is huge.
				go func(u string) {
					res := checker.CheckSite(u)
					db.UpdateResult(u, res)
					fmt.Printf("Checked %s: %s\n", u, res.Status)
				}(url)
			}
			// Hardcoded interval
			time.Sleep(10 * time.Second)
		}
	}()

	// 4. Start HTTP Server
	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		data := db.GetAll()
		
		// Manual JSON encoding? Nah, let's use the encoder
		if err := json.NewEncoder(w).Encode(data); err != nil {
			http.Error(w, "Failed to encode response", 500)
		}
	})

	port := ":8080"
	fmt.Printf("Server listening on port %s\n", port)
	// INTENTIONAL GAP: Ignoring errors from ListenAndServe
	http.ListenAndServe(port, nil)
}
