package main

import "fmt"

func main() {
	fmt.Println(ReduceFraction([2]int{60, 20}))
}

// First version
func ReduceFraction(fraction [2]int) [2]int {
	g := gcd(fraction[0], fraction[1])
	fraction[0] = fraction[0] / g
	fraction[1] = fraction[1] / g
	return fraction
}

func gcd(a, b int) int {
	for a > 0 && b > 0 {
		if a >= b {
			a = a % b
		} else {
			b = b % a
		}
	}

	if a > b {
		return a
	}
	return b
}

// https://www.codewars.com/kata/576400f2f716ca816d001614
//
// Write a function which reduces fractions to their simplest form!
// Fractions will be presented as an array/tuple (depending on the language)
// of strictly positive integers, and the reduced fraction must be returned as an array/tuple:

// input:   [numerator, denominator]
// output:  [reduced numerator, reduced denominator]
// example: [45, 120] --> [3, 8]
// All numerators and denominators will be positive integers.
