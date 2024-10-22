package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"io/ioutil"
	"log"
	"strings"
)

// Item represents an inventory item
type Item struct {
	ID    string  `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

// In-memory list of items
var items = []Item{
	{ID: "1", Name: "Laptop", Price: 799.99},
	{ID: "2", Name: "Smartphone", Price: 499.99},
}

// Get all items
func getItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}

// Get an item by ID
func getItemByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := strings.TrimPrefix(r.URL.Path, "/items/")
	for _, item := range items {
		if item.ID == id {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.Error(w, "Item not found", http.StatusNotFound)
}

// Create a new item
func createItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	var newItem Item
	if err := json.Unmarshal(body, &newItem); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	items = append(items, newItem)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(newItem)
}

// Update an item by ID
func updateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := strings.TrimPrefix(r.URL.Path, "/items/")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	var updatedItem Item
	if err := json.Unmarshal(body, &updatedItem); err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	for i, item := range items {
		if item.ID == id {
			items[i] = updatedItem
			json.NewEncoder(w).Encode(updatedItem)
			return
		}
	}
	http.Error(w, "Item not found", http.StatusNotFound)
}

// Delete an item by ID
func deleteItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	id := strings.TrimPrefix(r.URL.Path, "/items/")
	for i, item := range items {
		if item.ID == id {
			items = append(items[:i], items[i+1:]...)
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "Item with ID %s deleted", id)
			return
		}
	}
	http.Error(w, "Item not found", http.StatusNotFound)
}

// Router function to handle requests
func router(w http.ResponseWriter, r *http.Request) {
	switch {
	case r.Method == http.MethodGet && r.URL.Path == "/items":
		getItems(w, r)
	case r.Method == http.MethodGet && strings.HasPrefix(r.URL.Path, "/items/"):
		getItemByID(w, r)
	case r.Method == http.MethodPost && r.URL.Path == "/items":
		createItem(w, r)
	case r.Method == http.MethodPut && strings.HasPrefix(r.URL.Path, "/items/"):
		updateItem(w, r)
	case r.Method == http.MethodDelete && strings.HasPrefix(r.URL.Path, "/items/"):
		deleteItem(w, r)
	default:
		http.Error(w, "Not Found", http.StatusNotFound)
	}
}

func main() {
	// Set up routes
	http.HandleFunc("/items", router)
	http.HandleFunc("/items/", router)

	// Start the server on port 8080
	fmt.Println("Server is running on port 3000...")
	log.Fatal(http.ListenAndServe(":3000", nil))
}
