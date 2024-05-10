package main

import "fmt"

func main() {
	fmt.Println(ReverseList([]int{1, 2, 3, 4}))
}

func ReverseList(lst []int) (response []int) {
	response = make([]int, len(lst))
	for i, ii := 0, len(lst)-1; ii >= 0; i, ii = i+1, ii-1 {
		response[ii] = lst[i]
	}
	return response
}

// https://codewars.com/kata/57a04da9e298a7ee43000111
//
// Write reverseList function that simply reverses lists.
