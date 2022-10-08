package main

import "fmt"

func main() {
	fmt.Println(Between(1, 4))
}

func Between(a, b int) (response []int) {
	for i := a; i <= b; i++ {
		response = append(response, i)
	}
	return
}

// Complete the function that takes two integers (a, b, where a < b) and return
// an array of all integers between the input parameters, including them.
