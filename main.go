package main

import (
	"fmt"
	"go_lang/server"
)

func main() {
	name := "maharaja"
	details := &name
	consist := *details
	fmt.Println(details)
	fmt.Println(consist)

	dash()

	Map := make(map[string]int)
	Map["raja"] = 25

}

func dash() {
	variable := []string{"maga", "raja", "srinath"}
	for _, value := range variable {
		if value == variable[len(variable)-1] {
			fmt.Println(value)
			break
		}
		server.Server()
	}

}
