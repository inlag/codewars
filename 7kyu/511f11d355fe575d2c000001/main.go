package main

import (
	"fmt"
)

func main() {
	fmt.Println(TwoOldestAges([]int{6, 5, 83, 5, 3, 18}))
	fmt.Println(TwoOldestAges([]int{39, 53, 83, 51, 59, 61, 95, 23, 99, 49}))
	fmt.Println(TwoOldestAges([]int{93, 35, 53, 67, 17, 23, 89, 75, 15, 53}))
}

func TwoOldestAges(ages []int) [2]int {
	var one, two = 0, 0
	for i := 1; i < len(ages); i++ {
		if ages[i] > ages[two] {
			one = two
			two = i
		} else if two == one {
			one = i
		}
		if i != two && ages[i] > ages[one] {
			one = i
		}
	}
	return [2]int{ages[one], ages[two]}
}

// https://www.codewars.com/kata/511f11d355fe575d2c000001
//
//
//The two oldest ages function/method needs to be completed.
// It should take an array of numbers as its argument and return the two highest numbers within the array.
// The returned value should be an array in the format [second oldest age,  oldest age].

// The order of the numbers passed in could be any order. The array will always include at least 2 items.
// If there are two or more oldest age, then return both of them in array format.

// For example (Input --> Output):

// [1, 2, 10, 8] --> [8, 10]
// [1, 5, 87, 45, 8, 8] --> [45, 87]
// [1, 3, 10, 0]) --> [3, 10]
