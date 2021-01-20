package main
import (
	"fmt"
	"math/rand"
)

func main() {
	// var count = 0
	// for count < 10 {
	// 	var num = rand.Intn(10) + 1
	// 	fmt.Println(num)
	// 	count++
	// }
	// // can't access num

	// `count := 10` is equivalent to `var count = 0`
	for count := 10; count > 0; count-- {
		fmt.Println(count)
	}

	// if num := rand.Intn(3); num == 0 {
	// 	fmt.Println("Space Adventures")
	// } else if num == 1 {
	// 	fmt.Println("SpaceX")
	// } else {
	// 	fmt.Println("Virgin Galactic")
	// }
	switch num := rand.Intn(10); num {
	case 0:
		fmt.Println("Space Adventures")
	case 1:
		fmt.Println("SpaceX")
	case 2:
		fmt.Println("Virgin Galactic")
	default:
		fmt.Println("Random spaceline #", num)
	}
}