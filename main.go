package main

import (
	"fmt"
)

func main() {
	name := "hello_world"
	var nude string = "ramanujam"
	var number int = 43

	fmt.Println(number)
	fmt.Println(name)
	fmt.Println(nude)

	var array [3]string = [3]string{"maga", "raja", "karthi"}

	data := array[:]

	data = append(data, "king", "quen")
	fmt.Println(data)

	dev()
}

func dev() {

	//_nepal := 56

	// myMap := make(map[string]int)

	// myMap["magaraja"] = 2786745323456775
	// myMap["magja"] = 57

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// fmt.Println(myMap["magaraja"])

	names := "badass"

	devops(names)
}

func devops(name string) {

	fmt.Println(name)
}
