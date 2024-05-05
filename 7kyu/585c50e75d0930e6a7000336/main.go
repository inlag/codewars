package main

import (
	"fmt"
)

func main() {
	fmt.Println(AreCoprime(20, 27))
	fmt.Println(AreCoprime(12, 39))
}

func AreCoprime(n, m int) bool {
	if gcd(n, m) == 1 {
		return true
	}
	return false
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

// https://www.codewars.com/kata/585c50e75d0930e6a7000336
//
// Write a program to determine if the two given numbers are coprime.
// A pair of numbers are coprime if their greatest shared factor is 1.

// The inputs will always be two positive integers between 2 and 99.

// Examples
// 20 and 27:

// Factors of 20: 1, 2, 4, 5, 10, 20
// Factors of 27: 1, 3, 9, 27
// Greatest shared factor: 1
// Result: 20 and 27 are coprime
// 12 and 39:

// Factors of 12: 1, 2, 3, 4, 6, 12
// Factors of 39: 1, 3, 13, 39
// Greatest shared factor: 3
// Result: 12 and 39 are not coprimes
