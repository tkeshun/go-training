package main

import (
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"golang.org/x/sync/errgroup"
)

func main() {
	// Base URL of the mock server
	baseURL := "http://localhost:8080/api/"

	startTime := time.Now()

	// Map to store responses
	responses := make(map[int]string)
	var mu sync.Mutex

	var g errgroup.Group

	for i := 0; i < 20; i++ {
		// Capture the current value of i to avoid goroutine issues
		i := i
		url := fmt.Sprintf("%s%d", baseURL, i)

		g.Go(func() error {
			// Make HTTP GET request
			resp, err := http.Get(url)
			if err != nil {
				return fmt.Errorf("error fetching %s: %v", url, err)
			}
			defer resp.Body.Close()

			// Read the response body
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				return fmt.Errorf("error reading response from %s: %v", url, err)
			}

			// Store the response in the map
			mu.Lock()
			responses[i] = string(body)
			mu.Unlock()
			return nil
		})
	}

	if err := g.Wait(); err != nil {
		fmt.Printf("Error: %v\n", err)
	}

	for i := 0; i < 20; i++ {
		fmt.Printf("Response from /api/%d: %s\n", i, responses[i])
	}

	elapsedTime := time.Since(startTime)
	fmt.Printf("Total elapsed time: %v\n", elapsedTime)
}
