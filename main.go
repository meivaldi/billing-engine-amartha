package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// Define a struct for the response
type Response struct {
	Message string `json:"message"`
}

func main() {
	// Define the handler for the root endpoint
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		response := Response{Message: "Billing Engine"}
		jsonResponse(w, response)
	})

	// Define the handler for the /greet endpoint
	http.HandleFunc("/greet", func(w http.ResponseWriter, r *http.Request) {
		name := r.URL.Query().Get("name")
		if name == "" {
			name = "Guest"
		}
		response := Response{Message: fmt.Sprintf("Hello, %s!", name)}
		jsonResponse(w, response)
	})

	// Start the server
	fmt.Println("Starting server on port 8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// jsonResponse writes a JSON response to the ResponseWriter
func jsonResponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
