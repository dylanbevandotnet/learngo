package main

import (
	"fmt"
)

// how long does it take to get to Mars
func main() {
	const lightSpeed = 299792 // km/s

	// planets closest to each other
	var distance = 56000000 // km
	fmt.Println(distance/lightSpeed, "seconds")

	// opposite ends of the sun
	distance = 401000000
	fmt.Println(distance/lightSpeed, "seconds")

	// declare multiple variables
	var (
		closeDistance = 56000000
		farDistance = 401000000
	)

	//var closeDistance, farDistance = 56000000, 401000000
	fmt.Println(closeDistance/lightSpeed, "seconds")
	fmt.Println(farDistance/lightSpeed, "seconds")

	var weight = 149.0
	// weight = weight * 0.3783
	weight *= 0.3783
	fmt.Println("Weight is ", weight)

	var age = 41
	// age = age + 1
	// age += 1
	age++
	fmt.Println("Age is ", age)
}