package main

import "fmt"

func main() {
	fmt.Println(ZeroFuel(50, 25, 2)) // true
	fmt.Println(ZeroFuel(70, 25, 1)) // false
}

func ZeroFuel(distanceToPump int, mpg int, fuelLeft int) bool {
	return distanceToPump <= mpg*fuelLeft
}
