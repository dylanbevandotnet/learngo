package main

import (
	"fmt"
)

// Malacandra is the name for Mars in C.S. Lewis' "The Space Trilogy"
func main() {
	// determine how fast a ship would need to travel in km/h to reach Malacandra in 28 days.
	// Assume a distance of 56,000,000 km

	const (
		distance = 56000000
		totalFlightTimeInHours = 28 * 24
	)

	fmt.Printf("You must travel at %v km/h to get to Malacandra in 28 days\n",  distance/totalFlightTimeInHours)
}