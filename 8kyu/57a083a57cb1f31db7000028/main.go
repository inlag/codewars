package main

import (
	"fmt"
)

func main() {
	fmt.Println(PowersOfTwo(0))
	fmt.Println(PowersOfTwo(1))
	fmt.Println(PowersOfTwo(4))
}

func PowersOfTwo(n int) []uint64 {
	var response = make([]uint64, 0)
	for i := 0; i <= n; i++ {
		response = append(response, uint64(recursiveExponents(2, i)))
	}
	return response
}

func recursiveExponents(number, exponent int) int {
	if exponent == 0 {
		return 1
	}
	return number * recursiveExponents(number, exponent-1)
}

//Complete the function that takes a non-negative integer n as input, and
//returns a list of all the powers of 2 with the exponent ranging from 0 to n (
//inclusive ).
