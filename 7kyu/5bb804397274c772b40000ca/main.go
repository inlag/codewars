package main

import (
	"fmt"
	"math"
)

func StackHeight2d(layers int) float64 {
	if layers == 0 {
		return 0.000
	}
	return float64((layers-1))*math.Sqrt(float64(3))/2 + 1
}
func main() {
	fmt.Println(StackHeight2d(2))
}

// https://www.codewars.com/kata/5bb804397274c772b40000ca
//
//Background
// I have stacked some pool balls in a triangle.

// Like this,

// pool balls
// Kata Task
// Given the number of layers of my stack, what is the total height?

// Return the height as multiple of the ball diameter.

// Example
// The image above shows a stack of 5 layers.

// Notes
// layers >= 0
// approximate answers (within 0.001) are good enough
