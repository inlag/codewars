package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(NearestSq(2))    // 1
	fmt.Println(NearestSq(10))   // 9
	fmt.Println(NearestSq(111))  // 121
	fmt.Println(NearestSq(9999)) // 10000
}

func NearestSq(n int) int {
	return int(math.Pow(math.Round(math.Sqrt(float64(n))), 2))
}

// https://www.codewars.com/kata/5a805d8cafa10f8b930005ba
// Your task is to find the nearest square number, nearest_sq(n) or nearestSq(n), of a positive integer n.
//
// For example, if n = 111, then nearest\_sq(n) (nearestSq(n)) equals 121, since 111 is closer to 121, the square of 11, than 100, the square of 10.
//
// If the n is already the perfect square (e.g. n = 144, n = 81, etc.), you need to just return n.
//
// Good luck :)
//
// Check my other katas:
//
// Alphabetically ordered
//
// Case-sensitive!
//
// Not prime numbers
//
// Find your caterer
