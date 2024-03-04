package main

import "fmt"

func main() {
	fmt.Println(Gcd(30, 12)) // 6
	fmt.Println(Gcd(8, 9))   // 1
	fmt.Println(Gcd(1, 1))   // 1
}

func Gcd(x, y uint32) uint32 {
	for y != 0 {
		t := y
		y = x % y
		x = t
	}
	return x
}

// https://www.codewars.com/kata/5500d54c2ebe0a8e8a0003fd
//
// Find the greatest common divisor of two positive integers. The integers can
// be large, so you need to find a clever solution.
//
// The inputs x and y are always greater or equal to 1, so the greatest common
// divisor will always be an integer that is also greater or equal to 1.
