package main
import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Printf("%-16v %4v %-10v %5v\n", "Spaceline", "Days", "Trip type", "Price")
	fmt.Println("========================================")

	for ticketNum := 0; ticketNum < 10; ticketNum++ {
		const distance = 62100000
		var (
			spaceLine = ""
			speed = rand.Intn(15) + 16
			// departureDate = "October 13th 2020"
			roundTrip = "One-way"
		)

		// calculate the duration
		duration := (distance/speed)/(60*60*24)

		// generate the spaceline
		switch num := rand.Intn(3); num {
		case 0:
			spaceLine = "SpaceX"
		case 1:
			spaceLine = "Space Adventures"
		case 2:
			spaceLine = "Virgin Galactic"
		}

		// calculate the cost
		cost := (speed + 20)

		switch num := rand.Intn(2); num {
		case 0:
			roundTrip = "Round-trip"
			cost*=2
		}

		fmt.Printf("%-16v %4v %-10v $%4v\n", spaceLine, duration, roundTrip, cost)
	}
}