package main

import "fmt"

func main() {
	fmt.Println(SmallestIntegerFinder([]int{34, 15, 88, 2})) // 2
}

func SmallestIntegerFinder(numbers []int) int {
	var res int = numbers[0]
	for _, number := range numbers[0:] {
		if res > number {
			res = number
		}
	}
	return res
}

//Given an array of integers your solution should find the smallest integer.
//
//For example:
//
//Given [34, 15, 88, 2] your solution will return 2
//Given [34, -345, -1, 100] your solution will return -345
//You can assume, for the purpose of this kata, that the supplied array will not be empty.
