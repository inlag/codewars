package main

import "fmt"

func main() {
	fmt.Println(CountBy(50, 5))
}

func CountBy(x, n int) (response []int) {
	var i = 0
	for len(response) != n {
		i++
		if i%x == 0 {
			response = append(response, i)
		}
	}
	return
}

//Create a function with two arguments that will return an array of the first n
//multiples of x.
//
//Assume both the given number and the number of times to count will be positive
//numbers greater than 0.
//
//Return the results as an array or list ( depending on language ).
