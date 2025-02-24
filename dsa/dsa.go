package main

import "fmt"

// Insert an element at a specific index
func insert(arr []int, index, value int) []int {
	if index < 0 || index > len(arr) {
		fmt.Println("Invalid index")
		return arr
	}
	arr = append(arr[:index], append([]int{value}, arr[index:]...)...)
	return arr
}

// Delete an element at a specific index
func delete(arr []int, index int) []int {
	if index < 0 || index >= len(arr) {
		fmt.Println("Invalid index")
		return arr
	}
	arr = append(arr[:index], arr[index+1:]...)
	return arr
}

// Search for an element and return its index
func search(arr []int, value int) int {
	for i, v := range arr {
		if v == value {
			return i
		}
	}
	return -1 // Return -1 if value not found
}

// Traverse and print all elements
func traverse(arr []int) {
	fmt.Println("Array elements:")
	for i, v := range arr {
		fmt.Printf("Index %d: %d\n", i, v)
	}
}

func main() {
	// Initialize an array
	arr := []int{10, 20, 30, 40, 50}
	fmt.Println("Initial Array:", arr)

	// Insert element
	arr = insert(arr, 2, 25)
	fmt.Println("After Insertion:", arr)

	// Delete element
	arr = delete(arr, 4)
	fmt.Println("After Deletion:", arr)

	// Search for an element
	index := search(arr, 25)
	if index != -1 {
		fmt.Printf("Element 25 found at index %d\n", index)
	} else {
		fmt.Println("Element 25 not found")
	}

	// Traverse the array
	traverse(arr)
}
