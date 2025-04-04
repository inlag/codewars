package main

import "fmt"

func main() {
	fmt.Println(SimpleMultiplication(8))
	fmt.Println(SimpleMultiplication(9))
}

func SimpleMultiplication(n int) int {
	if n%2 == 0 {
		return n * 8
	} else {
		return n * 9
	}
}
