package main

import (
	"fmt"
)

func main() {
	// fmt.Print("My Weight on the surface of Mars is ")
	// fmt.Print(149.0 * 0.3783)
	// fmt.Print(" lbs, and I would be ")
	// fmt.Print(41 * 365 / 687)
	// fmt.Print(" years old.")

	// fmt.Print("My Weight on the surface of Mars is ", 149.0 * 0.3783, " lbs, and I would be ", 41 * 365 / 687, " years old.")

	fmt.Printf("My Weight on the surface of Mars is %v lbs, and I would be %v years old.\n", 149.0 * 0.3783, 41 * 365 / 687)

	// use padding with %
	fmt.Printf("%-15v $%4v\n", "SpaceX", 94)
	fmt.Printf("%-15v $%4v\n", "Virgin Galactic", 100)
}