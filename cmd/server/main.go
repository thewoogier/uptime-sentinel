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
	// TODO: Fix the path being hardcoded relative to where the binary is run.
	targets, err := config.LoadTargets("./configs/targets.json")
	if err != nil {
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
				// TODO: Fix launching a goroutine for every check without a worker pool?
				go func(u string) {
					res := checker.CheckSite(u)
					db.UpdateResult(u, res)
					fmt.Printf("Checked %s: %s\n", u, res.Status)
				}(url)
			}
			time.Sleep(10 * time.Second)
		}
	}()

	// 4. Start HTTP Server
	http.HandleFunc("/status", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		data := db.GetAll()

		if err := json.NewEncoder(w).Encode(data); err != nil {
			http.Error(w, "Failed to encode response", 500)
		}
	})

	port := ":8080"
	fmt.Printf("Server listening on port %s\n", port)
	// TODO: Better error handling from ListenAndServe
	http.ListenAndServe(port, nil)
}
