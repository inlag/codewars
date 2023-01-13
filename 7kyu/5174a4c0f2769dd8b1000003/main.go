package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(SortNumbers([]int{1, 2, 10, 50, 5}))
	fmt.Println(SortNumbers([]int{}))
}

func SortNumbers(numbers []int) []int {
	sort.Ints(numbers)
	return numbers
}

//Finish the solution so that it sorts the passed in array of numbers. If the
//function passes in an empty array or null/nil value then it should return an
//empty array.
//
//For example:
//
//solution(c(1, 2, 3, 10, 5)) # should return c(1, 2, 3, 5, 10)
//solution(NULL)              # should return NULL
