package main

import "fmt"

func main() {
	fmt.Println(Invert([]int{1, 2, 3, 4, 5}))   // []int{-1,-2,-3,-4,-5}
	fmt.Println(Invert([]int{1, -2, 3, -4, 5})) // []int{-1, 2, -3, 4, -5}
	fmt.Println(Invert(nil))                    // nil
	fmt.Println(Invert([]int{0}))               // []int{0}
}

func Invert(arr []int) []int {
	var response = make([]int, len(arr), cap(arr))
	for i := 0; i < len(arr); i++ {
		response[i] = arr[i] * (-1)
	}
	return response
}

//Given a set of numbers, return the additive inverse of each. Each positive becomes negatives, and the negatives become positives.
//Don't mutate the input!
//invert([1,2,3,4,5]) == [-1,-2,-3,-4,-5]
//invert([1,-2,3,-4,5]) == [-1,2,-3,4,-5]
//invert([]) == []
