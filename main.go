package main

import (
	"fmt"
)

func main() {
	students := []string{"jhon", "nithis", "mairandi"}
	numbers := []int{1, 2, 3, 4, 5, 6}

	for _, name := range students {
		fmt.Println(name)
	}
	fmt.Println(numbers)
}
