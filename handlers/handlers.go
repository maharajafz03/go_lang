// handlers/handlers.go
package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// HomeHandler responds with a welcome message
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to the homepage!")
}

// APIHandler responds with a JSON message
func APIHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "Hello from the API"}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// DanceHandler responds with a simple text message
func DanceHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprintln(w, "hello _da")
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse JSON body
	var data map[string]interface{}
	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	// Process data (for demonstration, we're just echoing back)
	response := map[string]string{
		"status": "success",
		"data":   fmt.Sprintf("%v", data),
	}

	// Set response header and encode response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
