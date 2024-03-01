package main

import "fmt"

func main() {
	fmt.Println(ArrowArea(4, 2))   // 2.0
	fmt.Println(ArrowArea(7, 6))   // 10.5
	fmt.Println(ArrowArea(25, 25)) //156.25
}

func ArrowArea(a, b int) float64 {
	return float64(a) * (float64(b) / 4)
}

// https://www.codewars.com/kata/589478160c0f8a40870000bc
//
//Area of an arrow
//An arrow is formed in a rectangle with sides a and b by joining the bottom corners to the midpoint of the top edge and the centre of the rectangle.
//
//arrow
//
//a and b are integers and > 0
//
//Write a function which returns the area of the arrow.
