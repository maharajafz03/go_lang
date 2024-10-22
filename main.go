package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// Set up routes
	http.HandleFunc("/items", router)
	http.HandleFunc("/items/", router)

	// Start the server on port 3000
	fmt.Println("Server is running on port 3000...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
