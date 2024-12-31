package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Response struct {
	Message   string `json:"message"`
	DelayTime int    `json:"delayTime"`
}

func main() {
	// Mux router to handle endpoints
	mux := http.NewServeMux()

	delayTime := []int{3, 4, 2, 5, 3, 1, 6, 3, 2, 4, 3, 1, 2, 5, 3, 2, 3, 4, 2, 2}

	// Define 20 endpoints
	for i, d := range delayTime {
		endpoint := fmt.Sprintf("/api/%d", i) // Create endpoint name
		mux.HandleFunc(endpoint, createHandler(i, d))
	}

	// Start server
	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", mux)
}

func createHandler(id int, delayTime int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Sleep for 3 seconds to simulate delay
		time.Sleep(time.Duration(delayTime) * time.Second)

		// Create response JSON
		response := Response{
			Message:   fmt.Sprintf("Endpoint %d OK", id),
			DelayTime: delayTime,
		}

		// Write JSON response
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
	}
}
