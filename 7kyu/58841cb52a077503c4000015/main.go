package main

import (
	"fmt"
)

func main() {
	fmt.Println(CircleOfNumbers(10, 2))
	fmt.Println(CircleOfNumbers(4, 1))
}

func CircleOfNumbers(n int, firstNumber int) int {
	halfN := n / 2
	num := firstNumber - halfN
	if num < 0 {
		return firstNumber + halfN
	}
	return num
}
