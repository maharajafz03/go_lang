package main

import "fmt"

func main() {

	// var inslice [3]int32 = make([3]int32, 3, 8, 5)

	numbers := make([]int, 5, 10) // Creates a slice of length 5 and capacity 10
	fmt.Println(numbers)
}
