package checker

import (
	"fmt"
	"net/http"
	"time"
)

type CheckResult struct {
	URL       string    `json:"url"`
	Status    string    `json:"status"`
	Latency   string    `json:"latency"`
	Timestamp time.Time `json:"timestamp"`
}

// CheckSite performs a simple GET request to the target URL
// TODO: Add timeout support
func CheckSite(url string) CheckResult {
	start := time.Now()
	resp, err := http.Get(url)
	duration := time.Since(start)

	result := CheckResult{
		URL:       url,
		Timestamp: time.Now(),
		Latency:   fmt.Sprintf("%dms", duration.Milliseconds()),
	}

	if err != nil {
		result.Status = "DOWN"
		return result
	}

	// TODO: Make more robust
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		result.Status = "UP"
	} else {
		result.Status = fmt.Sprintf("DOWN (%d)", resp.StatusCode)
	}

	resp.Body.Close()
	return result
}
