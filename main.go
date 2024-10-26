package main

import (
	"fmt"
	"net/http"
	"yamuna/handlers"
)

// Struct to hold JSON response data
type Response struct {
	Message string `json:"message"`
}

// Home handler

func main() {
	// Set up routes
	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/api", handlers.APIHandler)
	http.HandleFunc("/danceing", handlers.DanceHandler)
	http.HandleFunc("/post", handlers.PostHandler)

	fmt.Println("Server is running on http://localhost:3080")
	http.ListenAndServe(":3080", nil)
}
