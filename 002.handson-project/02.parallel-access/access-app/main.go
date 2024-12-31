package main

import (
	"fmt"
	"io"
	"net/http"
	"time"
)

func main() {
	// Base URL of the mock server
	baseURL := "http://localhost:8080/api/"

	startTime := time.Now()

	// Loop through all 20 endpoints
	for i := 0; i < 20; i++ {
		url := fmt.Sprintf("%s%d", baseURL, i)
		fmt.Printf("Requesting: %s\n", url)

		// Make HTTP GET request
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("Error fetching %s: %v\n", url, err)
			continue
		}
		defer resp.Body.Close()

		// Read and print the response body
		body, err := io.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("Error reading response from %s: %v\n", url, err)
			continue
		}

		fmt.Printf("Response from %s: %s", url, body)
	}
	elapsedTime := time.Since(startTime)
	fmt.Printf("Total elapsed time: %v\n", elapsedTime)
}
