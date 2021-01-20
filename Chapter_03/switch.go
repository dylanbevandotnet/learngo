package main

import (
	"fmt"
)

func main() {
	fmt.Println("There is a cavern entrance here and a path to the east.")
	var command = "go inside"

	switch command {
	case "go east":
		fmt.Println("You head futher up the mountain")
	case "enter cave", "go inside":
		fmt.Println("You find yourself in a dimly lit cavern.")
	case "read sign":
		fmt.Println("The sign reads 'No Minors'.")
	default:
		fmt.Println("Didn't quite get that")
	}

	var room = "lake"
	switch {
	case room == "cave":
		fmt.Println("You find yourself in a dimly lit cavern.")
	case room == "lake":
		fmt.Println("the ice seem solid enough.")
		fallthrough
	case room == "underwater":
		fmt.Println("The water is freezing cold.")
	}
}