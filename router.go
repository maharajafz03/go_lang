package main

import (
	"net/http"
	"strings"
)

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
