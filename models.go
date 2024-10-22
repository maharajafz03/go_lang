package main

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
