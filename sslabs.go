package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"
)

const maxAttempts = 30 // approx 5 minutes
const statusInProgress = "IN_PROGRESS"
const statusReady = "READY"
const statusError = "ERROR"

func analyzeHost(domain string) (SSLResponse, error) {
	url := "https://api.ssllabs.com/api/v3/analyze?host=" + domain + "&all=done"

	// Poll interval starts at 5s, increases to 10s once analysis is in progress
	pollInterval := 5 * time.Second
	var lastStatus string

	for attempt := 1; attempt <= maxAttempts; attempt++ {
		resp, err := http.Get(url)
		if err != nil {
			return SSLResponse{}, err
		}

		if resp.StatusCode != http.StatusOK {
			resp.Body.Close()
			return SSLResponse{}, fmt.Errorf("unexpected HTTP status: %s", resp.Status)
		}

		var result SSLResponse
		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			resp.Body.Close()
			return SSLResponse{}, err
		}
		resp.Body.Close()

		// Print only when status changes
		if result.Status != lastStatus {
			fmt.Println("Status:", result.Status)
			lastStatus = result.Status
		} else if result.Status == statusInProgress {
			fmt.Println("Waiting for SSL Labs analysis to complete...")
		}

		// Poll interval starts at 5s, increases to 10s once analysis is in progress
		if result.Status == statusInProgress {
			pollInterval = 10 * time.Second
		}

		if result.Status == statusReady {
			return result, nil
		}

		if result.Status == statusError {
			if result.StatusMessage != "" {
				return SSLResponse{}, fmt.Errorf("SSL Labs error: %s", result.StatusMessage)
			}
			return SSLResponse{}, fmt.Errorf("SSL Labs returned ERROR status")
		}

		time.Sleep(pollInterval)
	}

	return SSLResponse{}, errors.New("SSL Labs analysis did not complete within the expected time")
}
