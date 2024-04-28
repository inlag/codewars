package main

import (
	"fmt"
	"math"
)

func LargestPower(n uint64) int {
	if n == 1 {
		return -1
	}
	var (
		wait bool    = false
		res  float64 = 0
	)
	for !wait {
		newRes := res + 1
		if uint64(math.Pow(3, newRes)) < n {
			res = newRes
		} else {
			wait = true
		}
	}
	return int(res)
}

func main() {
	fmt.Println(LargestPower(1))
	fmt.Println(LargestPower(3))
	fmt.Println(LargestPower(4))
}

// https://www.codewars.com/kata/57be674b93687de78c0001d9
// Given a positive integer N, return the largest integer k such that 3^k < N.

// For example,

// LargestPower(3) // returns 0
// LargestPower(4) // returns 1
// You may assume that the input to your function is always a positive integer.
