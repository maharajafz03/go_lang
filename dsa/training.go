package main

import "fmt"

func nb() {
	bikers := []string{"yamaha", "bmw", "ducati", "kawasaki"}

	for _, v := range bikers {
		fmt.Println(v)
	}
}
