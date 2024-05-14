package main

import (
	"fmt"
)

func main() {
	fmt.Println(SumCubes(0))
	fmt.Println(SumCubes(1))
	fmt.Println(SumCubes(2))
	fmt.Println(SumCubes(3))
	fmt.Println(SumCubes(4))
	fmt.Println(SumCubes(10))
	fmt.Println(SumCubes(123))
	fmt.Println(SumCubes(15925))
	fmt.Println(SumCubes(60149))
}
func SumCubes(n int) int {
	var sum = 0
	for i := 0; i <= n; i++ {
		sum += i * i * i
	}
	return sum
}

// https://www.codewars.com/kata/59a8570b570190d313000037
//
// Write a function that takes a positive integer n, sums all the cubed values from
// 1 to n (inclusive), and returns that sum.

// Assume that the input n will always be a positive integer.

// Examples: (Input --> output)

// 2 --> 9 (sum of the cubes of 1 and 2 is 1 + 8)
// 3 --> 36 (sum of the cubes of 1, 2, and 3 is 1 + 8 + 27)
