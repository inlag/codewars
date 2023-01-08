package main

import "fmt"

func main() {
	fmt.Println(EachCons([]int{}, 3))
	fmt.Println(EachCons([]int{3, 5, 8, 13}, 1))
	fmt.Println(EachCons([]int{3, 5, 8, 13}, 2))
	fmt.Println(EachCons([]int{3, 5, 8, 13}, 3))
}

func EachCons(arr []int, n int) [][]int {
	response := make([][]int, 0)
	for i := 0; i < len(arr)-n+1; i++ {
		response = append(response, arr[i:i+n])
	}
	return response
}

//Create a method each_cons that accepts a list and a number n, and returns
//cascading subsets of the list of size n, like so:
//
//each_cons([1,2,3,4], 2)
// => [[1,2], [2,3], [3,4]]
//
//each_cons([1,2,3,4], 3)
// => [[1,2,3],[2,3,4]]
